package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/image/math/fixed"

	"github.com/llgcode/draw2d"
	"github.com/mooncaker816/freetype/truetype"
	ifont "golang.org/x/image/font"
)

type ThreeFont struct {
	Glyphs map[string]ThreeGlyph `json:"glyphs"`
	ThreeStatics
}
type ThreeGlyph struct {
	AdvanceWidth int    `json:"ha"`
	Xmin         int    `json:"x_min"`
	Xmax         int    `json:"x_max"`
	Path         string `json:"o"`
}
type ThreeStatics struct {
	FamilyName         string         `json:"familyName"`
	Ascender           int            `json:"ascender"`
	Descender          int            `json:"descender"`
	UnderlinePosition  int            `json:"underlinePosition"`
	UnderlineThickness int            `json:"underlineThickness"`
	BoundingBox        map[string]int `json:"boundingBox"`
	Resolution         int            `json:"resolution"`
	//origininfo         map[string]string `json:"ha"`
	//cssfontweight      string            `json:"cssfontweight"`
	//cssfontstyle       string            `json:"cssfontstyle"`
}

var font *truetype.Font
var stat ThreeStatics

func init() {
	start := time.Now()
	// b, err := ioutil.ReadFile("../resource/AppleSymbols-Regular.ttf")
	b, err := ioutil.ReadFile("../resource/m.ttf")
	// b, err := ioutil.ReadFile("../resource/test.ttf")
	if err != nil {
		log.Fatal("can not read ttf!")
	}
	font, err = truetype.Parse(b)
	if err != nil {
		log.Fatal("can not parse ttf!", err)
	}
	var fontSize int32 = 1 //pixel
	// get the original bound box, as font.Bounds(scale) will divide by font.FUnitsPerEm()
	var scale = fixed.Int26_6(font.FUnitsPerEm() * fontSize)
	stat = ThreeStatics{
		// FamilyName         : ""
		// Ascender           : font.
		// Descender          int            `json:"descender"`
		// UnderlinePosition  int            `json:"underlinePosition"`
		// UnderlineThickness int            `json:"underlineThickness"`

		BoundingBox: map[string]int{
			"xMax": int(font.Bounds(scale).Max.X),
			"xMin": int(font.Bounds(scale).Min.X),
			"yMax": int(font.Bounds(scale).Max.Y),
			"yMin": int(font.Bounds(scale).Min.Y),
		},
		Resolution: 1000,
	}
	fmt.Println(time.Since(start))
	// fmt.Println(font.FamilyName, font.Ascender, font.Descender, font.UnderlinePosition, font.UnderlineThickness)
	// fmt.Println(font.BoundingBox)
	// for key, value := range font.Glyphs {
	// 	fmt.Println(key, ":", value)
	// }
}

func main() {
	http.HandleFunc("/getpath", getpath)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getpath(w http.ResponseWriter, req *http.Request) {
	log.Println("in")
	str := req.URL.Query().Get("str")
	if str == "" {
		fmt.Fprintln(w, "no input string, pls input something!\n")
		return
	}
	log.Println("haha")
	// n := utf8.RuneCount([]byte(str))
	// charsmap := make(map[string]ThreeGlyph, n)
	charsmap := make(map[string][]ThreeFont)

	fupe := fixed.Int26_6(font.FUnitsPerEm())

	// str = "䷀䷁䷂䷃䷄䷅䷆䷇䷈䷉䷊䷋䷌䷍䷎䷏䷐䷑䷒䷓䷔䷕䷖䷗䷘䷙䷚䷛䷜䷝䷞䷟䷠䷡䷢䷣䷤䷥䷦䷧䷨䷩䷪䷫䷬䷭䷮䷯䷰䷱䷲䷳䷴䷵䷶䷷䷸䷹䷺䷻䷼䷽䷾䷿"
	str = "𝄞🀢🀣🀤🀥🀏"
	for _, char := range str {
		log.Println(string(char))
		ix := font.Index(char)
		log.Println(ix)
		g := &truetype.GlyphBuf{}
		err := g.Load(font, fupe, ix, ifont.HintingNone)
		if err != nil {
			log.Fatalf("glyph load failed: %v", err)
		}

		// fmt.Fprintf(w, "%+v\n", g)
		// fmt.Fprintf(w, "%v points: %+v\n", char, g.Points)
		threefonts := make([]ThreeFont, 0)
		path := new(draw2d.Path)
		e0 := 0

		for i, e1 := range g.Ends {
			buildpath(path, g.Points[e0:e1], 0, 0)
			e0 = e1
			oneCharMap := make(map[string]ThreeGlyph)
			oneCharMap[string(char)] = ThreeGlyph{
				AdvanceWidth: int(g.AdvanceWidth),
				Xmin:         int(g.Bounds.Min.X),
				Xmax:         int(g.Bounds.Max.X),
				Path:         threePath(path) + "z ",
			}
			// buildpath(path, g.Points, 0, 0)
			// fmt.Fprintf(w, "%+v\n%s\n", *path, path.String())
			if i == len(g.Ends)-1 {
				threefonts = append(threefonts, ThreeFont{oneCharMap, stat})
			}
		}
		charsmap[string(char)] = threefonts
	}

	// newfont := ThreeFont{charsmap, threeFont.ThreeStatics}
	//log.Println(newfont)
	data, _ := json.MarshalIndent(charsmap, "", "	")
	// data, _ := json.Marshal(newfont)
	fmt.Fprintf(w, "%s\n", data)
	return
}

func buildpath(path draw2d.PathBuilder, ps []truetype.Point, dx, dy float64) {
	if len(ps) == 0 {
		return
	}
	startX, startY := pointToF64Point(ps[0])
	path.MoveTo(startX+dx, startY+dy)
	q0X, q0Y, on0 := startX, startY, true
	for _, p := range ps[1:] {
		qX, qY := pointToF64Point(p)
		on := p.Flags&0x01 != 0
		if on {
			if on0 {
				path.LineTo(qX+dx, qY+dy)
			} else {
				path.QuadCurveTo(q0X+dx, q0Y+dy, qX+dx, qY+dy)
			}
		} else {
			if on0 {
				// No-op.
			} else {
				midX := (q0X + qX) / 2
				midY := (q0Y + qY) / 2
				path.QuadCurveTo(q0X+dx, q0Y+dy, midX+dx, midY+dy)
			}
		}
		q0X, q0Y, on0 = qX, qY, on
	}
	// Close the curve.
	if on0 {
		path.LineTo(startX+dx, startY+dy)
	} else {
		path.QuadCurveTo(q0X+dx, q0Y+dy, startX+dx, startY+dy)
	}
}

func pointToF64Point(p truetype.Point) (x, y float64) {
	return fUnitsToFloat64(p.X) * 64, fUnitsToFloat64(p.Y) * 64
}

func fUnitsToFloat64(x fixed.Int26_6) float64 {
	scaled := x << 2
	return float64(scaled/256) + float64(scaled%256)/256.0
}

func threePath(p *draw2d.Path) string {
	s := ""
	j := 0
	for _, cmd := range p.Components {
		switch cmd {
		case draw2d.MoveToCmp:
			s += fmt.Sprintf("m %f %f ", p.Points[j], p.Points[j+1])
			j = j + 2
		case draw2d.LineToCmp:
			s += fmt.Sprintf("l %f %f ", p.Points[j], p.Points[j+1])
			j = j + 2
		case draw2d.QuadCurveToCmp:
			// s += fmt.Sprintf("q %f %f %f %f ", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3])
			s += fmt.Sprintf("q %f %f %f %f ", p.Points[j+2], p.Points[j+3], p.Points[j], p.Points[j+1])
			j = j + 4
		case draw2d.CubicCurveToCmp:
			// s += fmt.Sprintf("b %f %f %f %f %f %f ", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3], p.Points[j+4], p.Points[j+5])
			s += fmt.Sprintf("b %f %f %f %f %f %f ", p.Points[j+4], p.Points[j+5], p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3])
			j = j + 6
		case draw2d.ArcToCmp:
			// s += fmt.Sprintf("a %f %f %f %f %f %f ", p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3], p.Points[j+4], p.Points[j+5])
			s += fmt.Sprintf("a %f %f %f %f %f %f ", p.Points[j+4], p.Points[j+5], p.Points[j], p.Points[j+1], p.Points[j+2], p.Points[j+3])
			j = j + 6
		case draw2d.CloseCmp:
			s += "z "
		}
	}
	return s
}
