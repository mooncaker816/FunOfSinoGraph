package ichang

// Shengxiao Values  鼠牛虎兔龙蛇马羊猴鸡狗猪
const (
	SxShu Shengxiao = iota
	SxNiu
	SxHu
	SxTu
	SxLong
	SxShe
	SxMa
	SxYang
	SxHou
	SxJi
	SxGou
	SxZhu
	NotSx
)

var shengxiaos = [...]Shengxiao{SxShu, SxNiu, SxHu, SxTu, SxLong, SxShe, SxMa, SxYang, SxHou, SxJi, SxGou, SxZhu}

// Shengxiao 生肖
type Shengxiao uint8

var shengxiaostr = "鼠牛虎兔龙蛇马羊猴鸡狗猪"

func (s Shengxiao) String() string {
	if s >= NotSx {
		return ""
	}
	return shengxiaostr[s*3 : (s+1)*3]
}
