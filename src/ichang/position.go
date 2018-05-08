package ichang

// Pos 方位
type Pos uint8

// Pos values 东南西北中
// 巽四（东南） 离九（正南） 坤二（西南）
// 震三（正东） 五 （中央）  兑七（正西）
// 艮八（东北） 坎一（正北） 乾六（西北）
const (
	East Pos = iota
	West
	South
	North
	Center
	NorthEast
	SouthEast
	NorthWest
	SouthWest
	NotPos
)

var poses = [...]Pos{East, West, South, North, Center, NorthEast, SouthEast, NorthWest, SouthWest}

func (p Pos) String() string {
	switch p {
	case East:
		return "正东"
	case West:
		return "正西"
	case South:
		return "正南"
	case North:
		return "正北"
	case Center:
		return "中央"
	case NorthEast:
		return "东北"
	case SouthEast:
		return "东南"
	case NorthWest:
		return "西北"
	case SouthWest:
		return "西南"
	default:
		return ""
	}
}
