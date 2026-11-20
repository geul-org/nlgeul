//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestSignGraph — each candidate of a sign at a unit's edge is judged by where the sign stands and how many stand in a row: 쉼표 one; 쌍점 one at a word's end or alone, or between letters that are not both Latin letters or digits with no other sign next to it and not at a unit's head (EPI:입자, 줄:1; not a:b, C:\); 쌍반점 only at a word's end or alone; 물결표 at a word's end or alone, any number, or between two Hangul syllables; 빗금 between two Hangul syllables or at a unit's edge (after the letters, right before them, alone), not between other letters; 줄임표 and 줄표 any number anywhere; 샵 and 골뱅이 one with letters right after or alone; 태그 and 명칭 one with letters right after; 제목 at an input line's start with a space after; 가운뎃점 one (a sign next to it does not matter); 더하기 and 연결 one with no other sign right next to it; the numbers of edges returned
package parse

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
		{"쌍점", 5, rule.Mark{Run: 1, Facts: rule.MarkEnds}, true},
		{"쌍점", 5, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"쌍점", 5, rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideAlnum}, false},
		{"쌍점", 5, rule.Mark{Run: 2, Facts: rule.MarkEnds}, false},
		{"쌍점", 5, rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideHangul}, true},
		{"쌍점", 5, rule.Mark{Run: 1, Right: rule.SideHangul}, true},
		{"쌍점", 5, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideAlnum | rule.SideDigit}, true},
		{"쌍점", 5, rule.Mark{Run: 1, Left: rule.SideAlnum, Facts: rule.MarkBySign}, false},
		{"쌍점", 5, rule.Mark{Run: 1, Right: rule.SideHangul, Facts: rule.MarkLeads}, false},
		{"쌍반점", 2, rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideHangul}, false},
		{"쌍반점", 2, rule.Mark{Run: 1, Facts: rule.MarkEnds}, true},
		{"물결표", 2, rule.Mark{Run: 3, Facts: rule.MarkEnds}, true},
		{"물결표", 2, rule.Mark{Run: 1}, false},
		{"물결표", 2, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, true},
		{"물결표", 2, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideAlnum | rule.SideDigit}, false},
		{"물결표", 2, rule.Mark{Run: 1, Right: rule.SideHangul}, false},
		{"빗금", 2, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, true},
		{"빗금", 2, rule.Mark{Run: 1, Left: rule.SideAlnum, Right: rule.SideHangul}, false},
		{"빗금", 2, rule.Mark{Run: 1, Left: rule.SideHangul, Facts: rule.MarkEnds}, true},
		{"빗금", 2, rule.Mark{Run: 1, Left: rule.SideAlnum, Facts: rule.MarkEnds}, true},
		{"빗금", 2, rule.Mark{Run: 1, Right: rule.SideHangul, Facts: rule.MarkLeads}, true},
		{"빗금", 2, rule.Mark{Run: 1, Facts: rule.MarkAlone}, true},
		{"줄임표", 0, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, true},
		{"줄임표", 0, rule.Mark{Run: 2, Facts: rule.MarkEnds}, true},
		{"줄표", 0, rule.Mark{Run: 3, Facts: rule.MarkAlone}, true},
		{"줄표", 0, rule.Mark{Run: 1, Left: rule.SideHangul, Right: rule.SideHangul}, true},
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
