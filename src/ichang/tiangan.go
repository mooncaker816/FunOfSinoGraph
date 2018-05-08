package ichang

// Tiangan Values 甲乙丙丁戊己庚辛壬癸
const (
	TgJia Tiangan = iota
	TgYi
	TgBing
	TgDing
	TgWu
	TgJi
	TgGeng
	TgXin
	TgRen
	TgGui
	NotTg
)

var tiangans = [...]Tiangan{TgJia, TgYi, TgBing, TgDing, TgWu, TgJi, TgGeng, TgXin, TgRen, TgGui}

// Tiangan 五行
type Tiangan uint8

var tianganstr = "甲乙丙丁戊己庚辛壬癸"

func (t Tiangan) String() string {
	if t >= NotTg {
		return ""
	}
	return tianganstr[t*3 : (t+1)*3]
}

// 甲乙→木。 甲为阳木，乙为阴木。
// 丙丁→火。 丙为阳火，丁为阴火。
// 戊己→土。 戊为阳土，己为阴土。
// 庚辛→金。 庚为阳金，辛为阴金。
// 壬癸→水。 壬为阳水，癸为阴水。

// ToWuxing 天干对应的五行
func (t Tiangan) ToWuxing() Wuxing {
	if t >= NotTg {
		panic("invalid Tiangan")
	}
	return wuxings[t/2]
}

// IsYang 天干是否为阳
func (t Tiangan) IsYang() bool {
	if t >= NotTg {
		panic("invalid Tiangan")
	}
	return t%2 == 0
}
