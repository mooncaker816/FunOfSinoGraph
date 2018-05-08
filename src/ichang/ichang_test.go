package ichang

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTianganToWuxing(t *testing.T) {
	tt := []struct {
		in     Tiangan
		want   Wuxing
		isYang bool
	}{
		{TgJia, WxMu, true},
		{TgYi, WxMu, false},
		{TgBing, WxHuo, true},
		{TgDing, WxHuo, false},
		{TgWu, WxTu, true},
		{TgJi, WxTu, false},
		{TgGeng, WxJin, true},
		{TgXin, WxJin, false},
		{TgRen, WxShui, true},
		{TgGui, WxShui, false},
	}
	for _, tc := range tt {
		if wx := tc.in.ToWuxing(); wx != tc.want {
			t.Errorf("%v ToWuxing() got %v; want %v", tc.in, wx, tc.want)
		}
		if yy := tc.in.IsYang(); yy != tc.isYang {
			t.Errorf("%v IsYang() got %v; want %v", tc.in, yy, tc.isYang)
		}
	}
}

func TestWuxingToTiangans(t *testing.T) {
	tt := []struct {
		in   Wuxing
		want []Tiangan
	}{
		{WxMu, []Tiangan{TgJia, TgYi}},
		{WxHuo, []Tiangan{TgBing, TgDing}},
		{WxTu, []Tiangan{TgWu, TgJi}},
		{WxJin, []Tiangan{TgGeng, TgXin}},
		{WxShui, []Tiangan{TgRen, TgGui}},
	}
	for _, tc := range tt {
		if tg := tc.in.ToTiangans(); !reflect.DeepEqual(tg, tc.want) {
			t.Errorf("%v ToTiangans() got %v; want %v", tc.in, tg, tc.want)
		}
	}
}

func TestDizhiToWuxing(t *testing.T) {
	tt := []struct {
		in     Dizhi
		want   Wuxing
		isYang bool
	}{
		{DzZi, WxShui, true},
		{DzChou, WxTu, false},
		{DzYin, WxMu, true},
		{DzMao, WxMu, false},
		{DzChen, WxTu, true},
		{DzSi, WxHuo, false},
		{DzWu, WxHuo, true},
		{DzWei, WxTu, false},
		{DzShen, WxJin, true},
		{DzYou, WxJin, false},
		{DzXu, WxTu, true},
		{DzHai, WxShui, false},
	}
	for _, tc := range tt {
		if wx := tc.in.ToWuxing(); wx != tc.want {
			t.Errorf("%v ToWuxing() got %v; want %v", tc.in, wx, tc.want)
		}
		if yy := tc.in.IsYang(); yy != tc.isYang {
			t.Errorf("%v IsYang() got %v; want %v", tc.in, yy, tc.isYang)
		}
	}
}

func TestWuxingToDizhis(t *testing.T) {
	tt := []struct {
		in   Wuxing
		want []Dizhi
	}{
		{WxMu, []Dizhi{DzYin, DzMao}},
		{WxHuo, []Dizhi{DzSi, DzWu}},
		{WxTu, []Dizhi{DzChen, DzXu, DzChou, DzWei}},
		{WxJin, []Dizhi{DzShen, DzYou}},
		{WxShui, []Dizhi{DzHai, DzZi}},
	}
	for _, tc := range tt {
		if tg := tc.in.ToDizhis(); !reflect.DeepEqual(tg, tc.want) {
			t.Errorf("%v ToDizhis() got %v; want %v", tc.in, tg, tc.want)
		}
	}
}

func SampleBaGong(t *testing.T) {
	for _, gong := range gong8s {
		fmt.Println(gong)
	}
	// Output:
	// ䷀䷫䷠䷋䷓䷖䷢䷍
	// ䷜䷻䷂䷾䷰䷶䷣䷆
	// ䷳䷕䷙䷨䷥䷉䷼䷴
	// ䷲䷏䷧䷟䷭䷯䷛䷐
	// ䷸䷈䷤䷩䷘䷔䷚䷑
	// ䷝䷷䷱䷿䷃䷺䷅䷌
	// ䷁䷗䷒䷊䷡䷪䷄䷇
	// ䷹䷮䷬䷞䷦䷎䷽䷵
}

func TestBelows(t *testing.T) {
	tt := []Gong8{QianGong, DuiGong, LiGong, ZhenGong, XunGong, KanGong, GenGong, KunGong}
	for i, vs := range gua64s {
		for _, v := range vs {
			if gong := v.Belongs(); gong != tt[i] {
				t.Errorf("%v Belongs() got %v; want %v", v, gong, tt[i])
			}
		}
	}
}

func TestGetYueGuaShen(t *testing.T) {
	bg := XunWeiFeng
	bgplus := bg.BuildGua64Plus(nil)
	if bgplus.yueguashen != 5 {
		t.Errorf("%v setYueGuaShen() got %v; want %v", bg, bgplus.yueguashen, 5)
	}
}

func TestGetShiShen(t *testing.T) {
	bg := HuoDiJin
	bgplus := bg.BuildGua64Plus(nil)
	if bgplus.shishen != 4 {
		t.Errorf("%v GetShiShen() got %v; want %v", bg, bgplus.shishen, 4)
	}
}

func TestBuGua(t *testing.T) {
	fmt.Println(BuGua())
}
