package ichang

// Liuqin Values 父母、子孙、官鬼、妻财、兄弟
const (
	FuMu Liuqin = iota
	ZiSun
	GuanGui
	QiCai
	XiongDi
	NotLq
)

var liuqins = [...]Liuqin{FuMu, ZiSun, GuanGui, QiCai, XiongDi} //父母、子孙、官鬼、妻财、兄弟
var liuqinstr = "父母子孙官鬼妻财兄弟"

// Liuqin 六亲
type Liuqin uint8

func (l Liuqin) String() string {
	if l >= NotLq {
		return ""
	}
	return liuqinstr[l*6 : (l+1)*6]
}
