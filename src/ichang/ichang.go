package ichang

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"math/bits"
	"strings"
	"sync"
)

// GuaXiang 卦象: 本卦，变卦，综卦，错挂，互卦
type GuaXiang struct {
	BenGua  Gua64
	BianGua Gua64
	ZongGua Gua64
	CuoGua  Gua64
	HuGua   Gua64
	by      map[int]bool
}

// BuGua 卜卦
func BuGua() GuaXiang {
	var gx GuaXiang
	gx.by = make(map[int]bool)
	bianyao := uint8(0)
	for i := 0; i < 6; i++ {
		fmt.Printf("模拟占卜第%d爻中...\n", i+1)
		sum := gen3rands()
		switch sum {
		case 0: // 三个0，太阴，变爻
			bianyao |= 1 << uint(i)
			gx.by[i+1] = true
		case 1: // 一个1，少阳
			gx.BenGua |= 1 << uint(i)
		case 2: // 两个1，少阴
			// nop
		case 3: // 三个1，太阳，变爻
			gx.BenGua |= 1 << uint(i)
			bianyao |= 1 << uint(i)
			gx.by[i+1] = true
		}
		// time.Sleep(5 * time.Second)
	}
	// 变爻取反
	if bianyao > 0 {
		gx.BianGua = gx.BenGua ^ Gua64(bianyao)
	} else {
		gx.BianGua = NotGua64
	}
	// 颠倒
	gx.ZongGua = Gua64(bits.Reverse8(uint8(gx.BenGua)) >> 2)
	// 每一爻取反
	gx.CuoGua = gx.BenGua ^ (1<<6 - 1)
	// 3,4,5爻为上挂，2,3,4爻为下挂
	gx.HuGua = ((gx.BenGua &^ 0xe3) << 1) | ((gx.BenGua &^ 0xf1) >> 1)
	return gx
}

func gen3rands() int {
	var ret int
	ch := make(chan uint64)
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			result, _ := rand.Int(rand.Reader, big.NewInt(2))
			ch <- result.Uint64()
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		ret += int(v)
	}
	// fmt.Println(ret)
	return ret
}

func (gx GuaXiang) String() string {
	var b strings.Builder
	b.WriteString(gx.BenGua.BuildGua64Plus(gx.by).String())
	return b.String()
}
