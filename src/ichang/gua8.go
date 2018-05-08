package ichang

// Gua8 八卦
type Gua8 uint8

// Gua8 values
const (
	KunGua  Gua8 = iota //000 ☷坤地 乙{未巳卯}，癸{丑亥酉}
	ZhenGua             //001 ☳震雷 庚{子寅辰}，庚{午申戌}
	KanGua              //010 ☵坎水 戊{寅辰午}，戊{申戌子}
	DuiGua              //011 ☱兑泽 丁{巳卯丑}，丁{亥酉未}
	GenGua              //100 ☶艮山 丙{辰午申}，丙{戌子寅}
	LiGua               //101 ☲离火 己{卯丑亥}，己{酉未巳}
	XunGua              //110 ☴巽风 辛{丑亥酉}，辛{未巳卯}
	QianGua             //111 ☰乾天 甲{子寅辰}，壬{午申戊}
	NotGua8
)

var gua8s = []Gua8{KunGua, ZhenGua, KanGua, DuiGua, GenGua, LiGua, XunGua, QianGua}

func (g Gua8) tiangan(isUp bool) Tiangan {
	switch g {
	case KunGua:
		if isUp {
			return TgGui
		}
		return TgYi
	case ZhenGua:
		return TgGeng
	case KanGua:
		return TgWu
	case DuiGua:
		return TgDing
	case GenGua:
		return TgBing
	case LiGua:
		return TgJi
	case XunGua:
		return TgXin
	case QianGua:
		if isUp {
			return TgRen
		}
		return TgJia
	default:
		panic("invalid Gua8")
	}
}

// 乾在内子寅辰，乾在外午申戌；
// 巽在内丑亥酉，巽在外未巳卯；
// 坎在内寅辰午，坎在外申戌子；
// 离在内卯丑亥，离在外酉未巳；
// 艮在内辰午申，艮在外戌子寅；
// 兑在内巳卯丑，兑在外亥酉未；
// 震在内子寅辰，震在外午申戌；
// 坤在内未巳卯，坤在外丑亥酉。
func (g Gua8) dizhi(isUp bool) []Dizhi {
	switch g {
	case KunGua: //未巳卯,丑亥酉
		if isUp {
			return []Dizhi{DzChou, DzHai, DzYou}
		}
		return []Dizhi{DzWei, DzSi, DzMao}
	case ZhenGua, QianGua: //子寅辰,午申戌
		if isUp {
			return []Dizhi{DzWu, DzShen, DzXu}
		}
		return []Dizhi{DzZi, DzYin, DzChen}
	case KanGua: //寅辰午,申戌子
		if isUp {
			return []Dizhi{DzShen, DzXu, DzZi}
		}
		return []Dizhi{DzYin, DzChen, DzWu}
	case DuiGua: //巳卯丑,亥酉未
		if isUp {
			return []Dizhi{DzHai, DzYou, DzWei}
		}
		return []Dizhi{DzSi, DzMao, DzChou}
	case GenGua: //辰午申,戌子寅
		if isUp {
			return []Dizhi{DzXu, DzZi, DzYin}
		}
		return []Dizhi{DzChen, DzWu, DzShen}
	case LiGua: //卯丑亥,酉未巳
		if isUp {
			return []Dizhi{DzYou, DzWei, DzSi}
		}
		return []Dizhi{DzMao, DzChou, DzHai}
	case XunGua: //丑亥酉,未巳卯
		if isUp {
			return []Dizhi{DzWei, DzSi, DzMao}
		}
		return []Dizhi{DzChou, DzHai, DzYou}
	default:
		panic("invalid Gua8")
	}
}

var gua8str = "☷☳☵☱☶☲☴☰"

func (g Gua8) String() string {
	if g >= NotGua8 {
		return ""
	}
	return gua8str[g*3 : (g+1)*3]
}

// ToWuxing 八卦对应的五行
func (g Gua8) ToWuxing() Wuxing {
	switch g {
	case QianGua, DuiGua:
		return WxJin
	case KanGua:
		return WxShui
	case GenGua, KunGua:
		return WxTu
	case ZhenGua, XunGua:
		return WxMu
	case LiGua:
		return WxHuo
	default:
		panic("invalid Gua8")
	}
}
