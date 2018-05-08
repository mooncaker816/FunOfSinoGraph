package ichang

// Dizhi Values 寅卯辰巳午未申酉戌亥子丑
const (
	DzYin Dizhi = iota
	DzMao
	DzChen
	DzSi
	DzWu
	DzWei
	DzShen
	DzYou
	DzXu
	DzHai
	DzZi
	DzChou
	NotDz
)

var dizhis = [...]Dizhi{DzYin, DzMao, DzChen, DzSi, DzWu, DzWei, DzShen, DzYou, DzXu, DzHai, DzZi, DzChou}

// Dizhi 地支
type Dizhi uint8

var dizhistr = "寅卯辰巳午未申酉戌亥子丑"

func (d Dizhi) String() string {
	if d >= NotDz {
		return ""
	}
	return dizhistr[d*3 : (d+1)*3]
}

// 　　阳支：寅、辰、午、申、戌、子；
// 　　阴支：卯、巳、未、酉、亥、丑。

// IsYang 地支是否为阳支
func (d Dizhi) IsYang() bool {
	if d >= NotDz {
		panic("invalid Dizhi")
	}
	return d%2 == 0
}

// 木火土金水
// 寅卯——木 0,1--0
// 巳午——火 3,4--1
// 申酉——金 6,7--3
// 亥子——水 9,10--4
// 辰未戌丑--土 2,5,8,11--2

// ToWuxing 地支对应的五行
func (d Dizhi) ToWuxing() Wuxing {
	switch d {
	case DzYin, DzMao:
		return WxMu
	case DzSi, DzWu:
		return WxHuo
	case DzShen, DzYou:
		return WxJin
	case DzHai, DzZi:
		return WxShui
	case DzChen, DzXu, DzChou, DzWei:
		return WxTu
	default:
		panic("invalid Dizhi")
	}
}
