package ichang

// Wuxing Values 木火土金水
const (
	WxMu Wuxing = iota
	WxHuo
	WxTu
	WxJin
	WxShui
	NotWx
)

var wuxings = [...]Wuxing{WxMu, WxHuo, WxTu, WxJin, WxShui} //木火土金水

// Wuxing 五行
type Wuxing uint8

var wuxingstr = "木火土金水"

func (w Wuxing) String() string {
	if w >= NotWx {
		return ""
	}
	return wuxingstr[w*3 : (w+1)*3]
}

// ToTiangans 五行对应的天干
func (w Wuxing) ToTiangans() []Tiangan {
	if w >= NotWx {
		panic("invalid Wuxing")
	}
	offset := w * 2
	return tiangans[offset : offset+2]
}

// ToDizhis 五行对应的地支
func (w Wuxing) ToDizhis() []Dizhi {
	switch w {
	case WxJin:
		return []Dizhi{DzShen, DzYou}
	case WxMu:
		return []Dizhi{DzYin, DzMao}
	case WxShui:
		return []Dizhi{DzHai, DzZi}
	case WxHuo:
		return []Dizhi{DzSi, DzWu}
	case WxTu:
		return []Dizhi{DzChen, DzXu, DzChou, DzWei}
	default:
		panic("invalid Wuxing")
	}
}

// ToPos 五行对应的方位
func (w Wuxing) ToPos() Pos {
	switch w {
	case WxJin:
		return West
	case WxMu:
		return East
	case WxShui:
		return North
	case WxHuo:
		return South
	case WxTu:
		return Center
	default:
		panic("invalid Wuxing")
	}
}

// Ke 五行相克
func (w Wuxing) Ke() Wuxing {
	switch w {
	case WxJin:
		return WxMu
	case WxMu:
		return WxTu
	case WxTu:
		return WxShui
	case WxShui:
		return WxHuo
	case WxHuo:
		return WxJin
	default:
		panic("invalid wuxing")
	}
}

// Sheng 五行相生
func (w Wuxing) Sheng() Wuxing {
	switch w {
	case WxJin:
		return WxShui
	case WxMu:
		return WxHuo
	case WxTu:
		return WxJin
	case WxShui:
		return WxMu
	case WxHuo:
		return WxTu
	default:
		panic("invalid wuxing")
	}
}
