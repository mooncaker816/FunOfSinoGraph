package ichang

import (
	"fmt"
	"strconv"
	"strings"
)

// Gua64 六十四卦
type Gua64 uint16

// Gua64 Values
const (
	KunWeiDi        Gua64 = iota //000000 ䷁坤为地
	DiLeiFu                      //000001 ䷗地雷复
	DiShuiShi                    //000010 ䷆地水师
	DiZeLin                      //000011 ䷒地泽临
	DiShanQian                   //000100 ䷎地山谦
	DiHuoMingYi                  //000101 ䷣地火明夷
	DiFengSheng                  //000110 ䷭地风升
	DiTianTai                    //000111 ䷊地天泰
	LeiDiYu                      //001000 ䷏雷地豫
	ZhenWeiLei                   //001001 ䷲震为雷
	LeiShuiJie                   //001010 ䷧雷水解
	LeiZeGuiMei                  //001011 ䷵雷泽归妹
	LeiShanXiaoGuo               //001100 ䷽雷山小过
	LeiHuoFeng                   //001101 ䷶雷火丰
	LeiFengHeng                  //001110 ䷟雷风恒
	LeiTianDaZhuang              //001111 ䷡雷天大壮
	ShuiDiBi                     //010000 ䷇水地比
	ShuiLeiZhun                  //010001 ䷂水雷屯
	KanWeiShui                   //010010 ䷜坎为水
	ShuiZeJie                    //010011 ䷻水泽节
	ShuiShanJian                 //010100 ䷦水山蹇
	ShuiHuoJiJi                  //010101 ䷾水火既济
	ShuiFengJing                 //010110 ䷯水风井
	ShuiTianXu                   //010111 ䷄水天需
	ZeDiCui                      //011000 ䷬泽地萃
	ZeLeiSui                     //011001 ䷐泽雷随
	ZeShuiKun                    //011010 ䷮泽水困
	DuiWeiZe                     //011011 ䷹兑为泽
	ZeShanXian                   //011100 ䷞泽山咸
	ZeHuoGe                      //011101 ䷰泽火革
	ZeFengDaGuo                  //011110 ䷛泽风大过
	ZeTianQue                    //011111 ䷪泽天夬
	ShanDiBo                     //100000 ䷖山地剥
	ShanLeiYi                    //100001 ䷚山雷颐
	ShanShuiMeng                 //100010 ䷃山水蒙
	ShanZeSun                    //100011 ䷨山泽损
	GengWeiShan                  //100100 ䷳艮为山
	ShanHuoBen                   //100101 ䷕山火贲
	ShanFengGu                   //100110 ䷑山风蛊
	ShanTianDaChu                //100111 ䷙山天大畜
	HuoDiJin                     //101000 ䷢火地晋
	HuoLeiShiKe                  //101001 ䷔火雷噬嗑
	HuoShuiWeiJi                 //101010 ䷿火水未济
	HuoZeKui                     //101011 ䷥火泽睽
	HuoShanLv                    //101100 ䷷火山旅
	LiWeiHuo                     //101101 ䷝离为火
	HuoFengDing                  //101110 ䷱火风鼎
	HuoTianDaYou                 //101111 ䷍火天大有
	FengDiGuan                   //110000 ䷓风地观
	FengLeiYi                    //110001 ䷩风雷益
	FengShuiHuan                 //110010 ䷺风水涣
	FengZeZhongFu                //110011 ䷼风泽中孚
	FengShanJian                 //110100 ䷴风山渐
	FengHuoJiaRen                //110101 ䷤风火家人
	XunWeiFeng                   //110110 ䷸巽为风
	FengTianXiaoChu              //110111 ䷈风天小畜
	TianDiPi                     //111000 ䷋天地否
	TianLeiWuWang                //111001 ䷘天雷无妄
	TianShuiSong                 //111010 ䷅天水讼
	TianZeLv                     //111011 ䷉天泽履
	TianShanDun                  //111100 ䷠天山遁
	TianHuoTongRen               //111101 ䷌天火同人
	TianFengGou                  //111110 ䷫天风姤
	QianWeiTian                  //111111 ䷀乾为天
	NotGua64
)

// 乾一(金)	䷀乾为天	䷫天风姤	䷠天山遁	䷋天地否	䷓风地观	䷖山地剥	䷢火地晋	䷍火天大有
// 兑二(金)	䷹兑为泽	䷮泽水困	䷬泽地萃	䷞泽山咸	䷦水山蹇	䷎地山谦	䷽雷山小过	䷵雷泽归妹
// 离三(火)	䷝离为火	䷷火山旅	䷱火风鼎	䷿火水未济	䷃山水蒙	䷺风水涣    ䷅天水讼	䷌天火同人
// 震四(木)	䷲震为雷	䷏雷地豫	䷧雷水解	䷟雷风恒	䷭地风升	䷯水风井	䷛泽风大过	䷐泽雷随
// 巽五(木)	䷸巽为风	䷈风天小畜	䷤风火家人	䷩风雷益	䷘天雷无妄	䷔火雷噬嗑	䷚山雷颐	䷑山风蛊
// 坎六(水)	䷜坎为水	䷻水泽节	䷂水雷屯	䷾水火既济	䷰泽火革	䷶雷火丰	䷣地火明夷	䷆地水师
// 艮七(土)	䷳艮为山	䷕山火贲	䷙山天大畜	䷨山泽损	䷥火泽睽	䷉天泽履	䷼风泽中孚	䷴风山渐
// 坤八(土)	䷁坤为地	䷗地雷复	䷒地泽临	䷊地天泰	䷡雷天大壮	䷪泽天夬	䷄水天需	䷇水地比

var gua64s = [][]Gua64{
	{QianWeiTian, TianFengGou, TianShanDun, TianDiPi, FengDiGuan, ShanDiBo, HuoDiJin, HuoTianDaYou},
	{DuiWeiZe, ZeShuiKun, ZeDiCui, ZeShanXian, ShuiShanJian, DiShanQian, LeiShanXiaoGuo, LeiZeGuiMei},
	{LiWeiHuo, HuoShanLv, HuoFengDing, HuoShuiWeiJi, ShanShuiMeng, FengShuiHuan, TianShuiSong, TianHuoTongRen},
	{ZhenWeiLei, LeiDiYu, LeiShuiJie, LeiFengHeng, DiFengSheng, ShuiFengJing, ZeFengDaGuo, ZeLeiSui},
	{XunWeiFeng, FengTianXiaoChu, FengHuoJiaRen, FengLeiYi, TianLeiWuWang, HuoLeiShiKe, ShanLeiYi, ShanFengGu},
	{KanWeiShui, ShuiZeJie, ShuiLeiZhun, ShuiHuoJiJi, ZeHuoGe, LeiHuoFeng, DiHuoMingYi, DiShuiShi},
	{GengWeiShan, ShanHuoBen, ShanTianDaChu, ShanZeSun, HuoZeKui, TianZeLv, FengZeZhongFu, FengShanJian},
	{KunWeiDi, DiLeiFu, DiZeLin, DiTianTai, LeiTianDaZhuang, ZeTianQue, ShuiTianXu, ShuiDiBi},
}

func (g Gua64) splitToGua8() [2]Gua8 {
	var ret [2]Gua8
	ret[0] = Gua8(g & 0x7)
	ret[1] = Gua8((g >> 3) & 0x7)
	return ret
}

func (g Gua64) splitToYao() [6]Yao {
	var ret [6]Yao
	bg := g
	for i := 0; i < 6; i++ {
		ret[i] = Yao(bg & 1)
		bg >>= 1
	}
	return ret
}

// Belongs 64卦属于哪一宫
func (g Gua64) Belongs() Gong8 {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gua64s[i][j] == g {
				return Gong8(gua64s[i][0].splitToGua8()[0])
			}
		}
	}
	return NotGong8
}

// IsPrimary 判断是否为宫首
func (g Gua64) IsPrimary() bool {
	gs := g.splitToGua8()
	return gs[0] == gs[1]
}

// Gua64Plus 详细信息
type Gua64Plus struct {
	bg                  [2]Gua64
	by                  map[int]bool
	name                [2]string
	gong                Gong8
	upGua8              [2]Gua8
	downGua8            [2]Gua8
	yaos                [2][6]Yao
	tgs                 [2][6]Tiangan
	dzs                 [2][6]Dizhi
	wxs                 [2][6]Wuxing
	lqs                 [2][6]Liuqin
	shi, ying           int
	shishen, yueguashen int
}

// BuildGua64Plus 卦详细信息
func (g Gua64) BuildGua64Plus(by map[int]bool) *Gua64Plus {
	// fmt.Println("building ", g)
	gp := new(Gua64Plus)
	gp.bg[0] = g
	gp.name[0] = strings.Fields(gua64fullstr)[g]
	if by != nil {
		gp.bg[1] = g
		for k := range by {
			gp.bg[1] ^= 1 << uint(k-1)
		}
		gp.name[1] = strings.Fields(gua64fullstr)[gp.bg[1]]
	}

	gp.by = by

	gp.gong = g.Belongs() //分宫
	gp.setUpDown()        //拆分上下卦
	gp.setYao()           //拆分六爻
	gp.setTiangan()       //纳天干
	gp.setDizhi()         //纳地支
	gp.setWuxing()        //纳五行
	gp.setShiYing()       //纳世应
	gp.setLiuqin()        //纳六亲
	gp.setShiShen()       //纳世身
	gp.setYueGuaShen()    //纳月卦身
	return gp
}

func (gp *Gua64Plus) setUpDown() {
	for i := 0; i < 2; i++ {
		gs := gp.bg[i].splitToGua8()
		gp.downGua8[i] = gs[0]
		gp.upGua8[i] = gs[1]
		if len(gp.by) == 0 {
			break
		}
	}
}

func (gp *Gua64Plus) setYao() {
	for i := 0; i < 2; i++ {
		yaos := gp.bg[i].splitToYao()
		for j := 0; j < 6; j++ {
			gp.yaos[i][j] = yaos[j]
		}
		if len(gp.by) == 0 {
			break
		}
	}
}

// setTiangan 纳天干
func (gp *Gua64Plus) setTiangan() {
	for i := 0; i < 2; i++ {
		gp.tgs[i][0] = gp.downGua8[i].tiangan(false)
		gp.tgs[i][1] = gp.tgs[i][0]
		gp.tgs[i][2] = gp.tgs[i][0]
		gp.tgs[i][3] = gp.upGua8[i].tiangan(true)
		gp.tgs[i][4] = gp.tgs[i][3]
		gp.tgs[i][5] = gp.tgs[i][3]
		if len(gp.by) == 0 {
			break
		}
	}
}

// setDizhi 纳地支
func (gp *Gua64Plus) setDizhi() {
	for i := 0; i < 2; i++ {
		copy(gp.dzs[i][:3], gp.downGua8[i].dizhi(false))
		copy(gp.dzs[i][3:], gp.upGua8[i].dizhi(true))
		if len(gp.by) == 0 {
			break
		}
	}
}

// setWuxing 纳五行
func (gp *Gua64Plus) setWuxing() {
	for i := 0; i < 2; i++ {
		for j, dz := range gp.dzs[i] {
			gp.wxs[i][j] = dz.ToWuxing()
		}
		if len(gp.by) == 0 {
			break
		}
	}
}

// setShiYing 装世应
// 八卦之首世六当，以下初爻轮上扬； 游魂八卦四爻上，归魂八卦三爻详。
func (gp *Gua64Plus) setShiYing() {
	if gp.upGua8[0] == gp.downGua8[0] {
		gp.shi, gp.ying = 6, 3
		return
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gua64s[i][j] == gp.bg[0] {
				switch j {
				case 0:
					gp.shi, gp.ying = 6, 3
				case 1:
					gp.shi, gp.ying = 1, 4
				case 2:
					gp.shi, gp.ying = 2, 5
				case 3:
					gp.shi, gp.ying = 3, 6
				case 4:
					gp.shi, gp.ying = 4, 1
				case 5:
					gp.shi, gp.ying = 5, 2
				case 6:
					gp.shi, gp.ying = 4, 1
				case 7:
					gp.shi, gp.ying = 3, 6
				}
				return
			}
		}
	}
}

// setLiuqin 装六亲
// 乾兑金兄土父传，木财火鬼水子然。（乾兑两宫，八卦俱属金）
// 坎宫水兄火为财, 土鬼金父木子来。 (坎宫属水)
// 坤艮土兄火为父，木鬼水财金子路。（坤艮两宫，八卦俱属土）
// 离宫火兄水为鬼, 土子木父金财助。（离宫属火)
// 震巽木兄水父母，金鬼火子财是土。（震巽两宫，八卦俱属木）
func (gp *Gua64Plus) setLiuqin() {
	for j := 0; j < 2; j++ {
		switch gp.gong {
		case QianGong, DuiGong:
			for i, wx := range gp.wxs[j] {
				switch wx {
				case WxJin:
					gp.lqs[j][i] = XiongDi
				case WxTu:
					gp.lqs[j][i] = FuMu
				case WxMu:
					gp.lqs[j][i] = QiCai
				case WxHuo:
					gp.lqs[j][i] = GuanGui
				case WxShui:
					gp.lqs[j][i] = ZiSun
				}
			}
		case KanGong:
			for i, wx := range gp.wxs[j] {
				switch wx {
				case WxShui:
					gp.lqs[j][i] = XiongDi
				case WxHuo:
					gp.lqs[j][i] = QiCai
				case WxTu:
					gp.lqs[j][i] = GuanGui
				case WxJin:
					gp.lqs[j][i] = FuMu
				case WxMu:
					gp.lqs[j][i] = ZiSun
				}
			}
		case KunGong, GenGong:
			for i, wx := range gp.wxs[j] {
				switch wx {
				case WxTu:
					gp.lqs[j][i] = XiongDi
				case WxShui:
					gp.lqs[j][i] = QiCai
				case WxMu:
					gp.lqs[j][i] = GuanGui
				case WxHuo:
					gp.lqs[j][i] = FuMu
				case WxJin:
					gp.lqs[j][i] = ZiSun
				}
			}
		case LiGong:
			for i, wx := range gp.wxs[j] {
				switch wx {
				case WxHuo:
					gp.lqs[j][i] = XiongDi
				case WxJin:
					gp.lqs[j][i] = QiCai
				case WxShui:
					gp.lqs[j][i] = GuanGui
				case WxMu:
					gp.lqs[j][i] = FuMu
				case WxTu:
					gp.lqs[j][i] = ZiSun
				}
			}
		case ZhenGong, XunGong:
			for i, wx := range gp.wxs[j] {
				switch wx {
				case WxMu:
					gp.lqs[j][i] = XiongDi
				case WxTu:
					gp.lqs[j][i] = QiCai
				case WxJin:
					gp.lqs[j][i] = GuanGui
				case WxShui:
					gp.lqs[j][i] = FuMu
				case WxHuo:
					gp.lqs[j][i] = ZiSun
				}
			}
		}
		if len(gp.by) == 0 {
			break
		}
	}
}

// setShiShen 安世身
// 子午持世身居初，丑未持世身居二。
// 寅申持世身居三，卯酉持世身居四。
// 辰戌持世身居五，巳亥持世身居六
func (gp *Gua64Plus) setShiShen() {
	gp.shishen = int((gp.dzs[0][gp.shi-1]+2)%6) + 1
}

// setYueGuaShen 安月卦身（不一定有）
// 阴世则从午月起，阳世还从子月生；欲得识其卦中意，从初数至世方真。
func (gp *Gua64Plus) setYueGuaShen() {
	var target Dizhi
	if gp.yaos[0][gp.shi-1] == YinYao {
		target = (DzWu + Dizhi(gp.shi-1)) % 12
	} else {
		target = (DzZi + Dizhi(gp.shi-1)) % 12
	}
	for i, dz := range gp.dzs[0] {
		if dz == target {
			gp.yueguashen = i + 1
			return
		}
	}
	//没有月卦身 0
}

// 后天八卦 - 一数坎兮二数坤 三震四巽数中分 五寄中宫六乾是 七兑八艮九离门
// 巽四 离九 坤二
// 震三 五   兑七
// 艮八 坎一 乾六

// 先天八卦 - 乾一,兑二,离三,震四,巽五,坎六,艮七,坤八
// 兑二 乾一 巽五
// 离三      坎六
// 震四 坤八 艮七

// utf-8 : 一个字符占3个字节
var gua64str = "䷁䷗䷆䷒䷎䷣䷭䷊䷏䷲䷧䷵䷽䷶䷟䷡䷇䷂䷜䷻䷦䷾䷯䷄䷬䷐䷮䷹䷞䷰䷛䷪䷖䷚䷃䷨䷳䷕䷑䷙䷢䷔䷿䷥䷷䷝䷱䷍䷓䷩䷺䷼䷴䷤䷸䷈䷋䷘䷅䷉䷠䷌䷫䷀"
var gua64fullstr = `䷁坤为地
䷗地雷复
䷆地水师
䷒地泽临
䷎地山谦
䷣地火明夷
䷭地风升
䷊地天泰
䷏雷地豫
䷲震为雷
䷧雷水解
䷵雷泽归妹
䷽雷山小过
䷶雷火丰
䷟雷风恒
䷡雷天大壮
䷇水地比
䷂水雷屯
䷜坎为水
䷻水泽节
䷦水山蹇
䷾水火既济
䷯水风井
䷄水天需
䷬泽地萃
䷐泽雷随
䷮泽水困
䷹兑为泽
䷞泽山咸
䷰泽火革
䷛泽风大过
䷪泽天夬
䷖山地剥
䷚山雷颐
䷃山水蒙
䷨山泽损
䷳艮为山
䷕山火贲
䷑山风蛊
䷙山天大畜
䷢火地晋
䷔火雷噬嗑
䷿火水未济
䷥火泽睽
䷷火山旅
䷝离为火
䷱火风鼎
䷍火天大有
䷓风地观
䷩风雷益
䷺风水涣
䷼风泽中孚
䷴风山渐
䷤风火家人
䷸巽为风
䷈风天小畜
䷋天地否
䷘天雷无妄
䷅天水讼
䷉天泽履
䷠天山遁
䷌天火同人
䷫天风姤
䷀乾为天`

func (g Gua64) String() string {
	if g >= NotGua64 {
		return ""
	}
	// return strings.Fields(gua64fullstr)[g]
	return gua64str[g*3 : (g+1)*3]
}

func (gp Gua64Plus) String() string {
	var b strings.Builder
	b.WriteString("====" + gp.gong.String() + ":" + gp.name[0])
	if len(gp.by) > 0 {
		b.WriteString(">>>>>>>>> " + gp.name[1])
	}
	b.WriteString("====\n")
	// ll := b.Len()
	for i := 5; i >= 0; i-- {
		b.WriteString(gp.lqs[0][i].String())
		b.WriteString(" ")
		b.WriteString(gp.tgs[0][i].String())
		b.WriteString(gp.dzs[0][i].String())
		b.WriteString(" ")
		b.WriteString(gp.wxs[0][i].String())
		b.WriteString(" ")
		b.WriteString(gp.yaos[0][i].String())
		b.WriteString(" ")
		if i == gp.shi-1 {
			b.WriteString("世")
		} else if i == gp.ying-1 {
			b.WriteString("应")
		} else {
			b.WriteString("  ")
		}
		l := b.Len()
		fmt.Println(l)
		// b.WriteString(strings.Repeat(" ", b.Len()-ll))
		if len(gp.by) > 0 && gp.by[i+1] == true {
			b.WriteString("     * ")
			b.WriteString(gp.lqs[1][i].String())
			b.WriteString(" ")
			b.WriteString(gp.tgs[1][i].String())
			b.WriteString(gp.dzs[1][i].String())
			b.WriteString(" ")
			b.WriteString(gp.wxs[1][i].String())
			b.WriteString(" ")
			b.WriteString(gp.yaos[1][i].String())
			b.WriteString(" ")
		}
		b.WriteString("\n")
		// ll = b.Len()
	}
	b.WriteString("世身：" + strconv.Itoa(gp.shishen))
	b.WriteString("\n")
	if gp.yueguashen > 0 {
		b.WriteString("月卦身：" + strconv.Itoa(gp.yueguashen))
	} else {
		b.WriteString("月卦身：无")
	}
	b.WriteString("\n")
	return b.String()
}
