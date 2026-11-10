//ff:func feature=mini type=rule control=iteration dimension=1
//ff:what TestJudge — one constituent, several in the order of their verdicts, the last word as predicate only, the word divided at its particles or else into the pieces of its tail, 미정, and a failing graph
package mini

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestJudge(t *testing.T) {
	g, _ := newGraph()
	cases := []struct {
		word string
		last bool
		mark string
		want string
	}{
		{"부하가", false, "부하+가", "주어"},
		{"이것은", false, "이것+은", "관형어|주어|목적어"},
		{"문장을", false, "문장+을", "목적어|관형어"},
		{"학교에서는", false, "학교+에서+는", "주어|부사어"},
		{"먹지만", false, "먹+지만", "서술어"},
		{"크게", false, "크+게", "서술어|부사어"},
		{"작은", true, "작+은", "서술어"},
		{"학생이야", false, "학생+이야", "서술어|주어|목적어"},
		{"학생이야", true, "학생+이+야", "서술어"},
		{"문제였다", true, "문제+였+다", "서술어"},
		{"학생인", false, "학생+인", "관형어"},
		{"학교를", true, "학교를", "미정"},
		{"매우", false, "매우", "미정"},
	}
	for _, c := range cases {
		if marked, got, err := judge(g, readWord(c.word, c.last, &sentence{})); err != nil || got != c.want || marked != c.mark {
			t.Errorf("%s: %s %s (%v)", c.word, marked, got, err)
		}
	}
	bad := toulmin.NewGraph("bad")
	bad.Rule(func(toulmin.Context, toulmin.Specs) (bool, any) { panic("x") })
	if _, _, err := judge(bad, readWord("매우", false, &sentence{})); err == nil {
		t.Fatal("a panicking rule passed")
	}
}
