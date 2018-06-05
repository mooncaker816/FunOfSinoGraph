package ichang

import (
	"time"
)

// Shichen Values 子丑寅卯辰巳午未申酉戌亥
const (
	ScZi Shichen = iota
	ScChou
	ScYin
	ScMao
	ScChen
	ScSi
	ScWu
	ScWei
	ScShen
	ScYou
	ScXu
	ScHai
	NotSc
)

var shichens = [...]Shichen{ScZi, ScChou, ScYin, ScMao, ScChen, ScSi, ScWu, ScWei, ScShen, ScYou, ScXu, ScHai}

// Shichen 时辰
type Shichen uint8

var shichenstr = "子丑寅卯辰巳午未申酉戌亥"

func (s Shichen) String() string {
	if s >= NotSc {
		return ""
	}
	return shichenstr[s*3 : (s+1)*3]
}

// TimeToSc converts time.Time to Shichen
func TimeToSc(t time.Time) Shichen {
	return Shichen(((t.Hour() + 1) / 2) % 12)
}
