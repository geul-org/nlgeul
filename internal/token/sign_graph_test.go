//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestSignGraph — each candidate of a sign at a unit's edge is judged by where the sign stands and how many stand in a row: 쉼표 one; 쌍점 one at a word's end or alone; 물결표 at a word's end or alone, any number; 샵 and 골뱅이 one with letters right after or alone; 태그 and 명칭 one with letters right after; 제목 at an input line's start with a space after; 가운뎃점 one (a sign next to it does not matter); 더하기 and 연결 one with no other sign right next to it; the numbers of edges returned
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSignGraph(t *testing.T) {
	cases := []struct {
		name  string
		edges int
		m     rule.Mark
		stand bool
	}{
		{"쉼표", 1, rule.Mark{Run: 1}, true},
		{"쉼표", 1, rule.Mark{Run: 2}, false},
		{"쌍점", 2, rule.Mark{Run: 1, Facts: rule.MarkEnds}, true},
		{"쌍점", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"쌍점", 2, rule.Mark{Run: 1}, false},
		{"쌍점", 2, rule.Mark{Run: 2, Facts: rule.MarkEnds}, false},
		{"쌍반점", 2, rule.Mark{Run: 1, Facts: rule.MarkEnds}, true},
		{"물결표", 1, rule.Mark{Run: 3, Facts: rule.MarkEnds}, true},
		{"물결표", 1, rule.Mark{Run: 1}, false},
		{"샵", 2, rule.Mark{Run: 1, Facts: rule.MarkLeads}, true},
		{"샵", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"샵", 2, rule.Mark{Run: 1}, false},
		{"샵", 2, rule.Mark{Run: 2, Facts: rule.MarkLeads}, false},
		{"골뱅이", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"태그", 2, rule.Mark{Run: 1, Facts: rule.MarkLeads}, true},
		{"태그", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, false},
		{"명칭", 2, rule.Mark{Run: 1, Facts: rule.MarkLeads}, true},
		{"명칭", 2, rule.Mark{Run: 2, Facts: rule.MarkLeads}, false},
		{headingName, 2, rule.Mark{Run: 2, Facts: rule.MarkHeads | rule.MarkSpaced}, true},
		{headingName, 2, rule.Mark{Run: 1, Facts: rule.MarkSpaced}, false},
		{headingName, 2, rule.Mark{Run: 1, Facts: rule.MarkHeads}, false},
		{"물음표", 0, rule.Mark{Run: 3}, true},
		{"가운뎃점", 1, rule.Mark{Run: 1}, true},
		{"가운뎃점", 1, rule.Mark{Run: 2}, false},
		{"가운뎃점", 1, rule.Mark{Run: 1, Facts: rule.MarkBySign}, true},
		{"더하기", 2, rule.Mark{Run: 1}, true},
		{"더하기", 2, rule.Mark{Run: 2}, false},
		{"더하기", 2, rule.Mark{Run: 1, Facts: rule.MarkBySign}, false},
		{"연결", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"연결", 2, rule.Mark{Run: 1, Facts: rule.MarkBySign}, false},
	}
	for _, c := range cases {
		g, edges := signGraph(c.name)
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyMark, c.m)
		v, _ := evaluate(g, ctx, false)
		if edges != c.edges || (v > 0) != c.stand {
			t.Errorf("%s %+v: %d edges, verdict %.2f, want %d edges, stands %v", c.name, c.m, edges, v, c.edges, c.stand)
		}
	}
}
