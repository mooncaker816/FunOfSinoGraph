package ichang

// Yao 爻
type Yao uint8

var yaostr = "⚋⚊"

// Values of Yao
const (
	YinYao Yao = iota
	YangYao
)

func (y Yao) String() string {
	return yaostr[y*3 : (y+1)*3]
}
