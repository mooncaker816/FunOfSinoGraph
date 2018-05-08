package ichang

import (
	"strings"
)

// Gong8 八宫
type Gong8 uint8

// Vaules of Gong8
const (
	KunGong  Gong8 = iota //000 ☷坤地 乙{未巳卯}，癸{丑亥酉}
	ZhenGong              //001 ☳震雷 庚{子寅辰}，庚{午申戌}
	KanGong               //010 ☵坎水 戊{寅辰午}，戊{申戌子}
	DuiGong               //011 ☱兑泽 丁{巳卯丑}，丁{亥酉未}
	GenGong               //100 ☶艮山 丙{辰午申}，丙{戌子寅}
	LiGong                //101 ☲离火 己{卯丑亥}，己{酉未巳}
	XunGong               //110 ☴巽风 辛{丑亥酉}，辛{未巳卯}
	QianGong              //111 ☰乾天 甲{子寅辰}，壬{午申戊}
	NotGong8
)

var gong8s = []Gong8{QianGong, KanGong, GenGong, ZhenGong, XunGong, LiGong, KunGong, DuiGong}
var gongstr = "坤宫震宫坎宫兑宫艮宫离宫巽宫乾宫"

func (g Gong8) String() string {
	if g >= NotGong8 {
		return ""
	}
	return gongstr[g*6 : (g+1)*6]
}

// List 按宫列出依次变化的卦象
func (g Gong8) List() string {
	var b strings.Builder
	bg := Gua64(g<<3 | g)
	b.WriteString(bg.String())
	for i := 0; i < 7; i++ {
		if i < 5 {
			bg ^= 1 << uint(i)
		} else if i == 5 {
			bg ^= 1 << 3
		} else {
			bg ^= 1<<3 - 1
		}
		b.WriteString(bg.String())
	}
	return b.String()
}

// ToWuxing 宫对应的五行
func (g Gong8) ToWuxing() Wuxing {
	return Gua8(g).ToWuxing()
}
