//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSignContest — one stretch of one sign as a contest: a candidate for each graph that stands (샵 and 태그 for # with letters right after it, the heading named with its count and holding it as its depth); the letters with no name when none stands; beside the candidates when all of them are by what the sign does for the words after it (## is 제목2 and ##), not beside 샵
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSignContest(t *testing.T) {
	sc := scanOf("태그")
	cases := []struct {
		text string
		m    rule.Mark
		end  int
		want string
	}{
		{"?", rule.Mark{Run: 1}, 4, `"물음표"/stem=false/depth=0/end=4`},
		{"#", rule.Mark{Run: 1, Facts: rule.MarkLeads}, 2, `"샵"/stem=false/depth=0/end=2; "태그"/stem=false/depth=0/end=2`},
		{"@", rule.Mark{Run: 1, Facts: rule.MarkLeads}, 2, `"골뱅이"/stem=false/depth=0/end=2; "명칭"/stem=false/depth=0/end=2`},
		{"##", rule.Mark{Run: 2, Facts: rule.MarkHeads | rule.MarkSpaced}, 0, `"제목2"/stem=false/depth=2/end=0; ""/stem=true/depth=0/end=0`},
		{"#", rule.Mark{Run: 1, Facts: rule.MarkHeads | rule.MarkSpaced | rule.MarkAlone}, 0, `"제목1"/stem=false/depth=1/end=0; "샵"/stem=false/depth=0/end=0`},
		{"::", rule.Mark{Run: 2, Facts: rule.MarkEnds}, 0, `""/stem=true/depth=0/end=0`},
		{"가", rule.Mark{Run: 1}, 0, `""/stem=true/depth=0/end=0`},
	}
	for _, c := range cases {
		b := sc.signContest(c.text, c.m, c.end)
		if got := signBrief(b); b.Text != c.text || got != c.want {
			t.Errorf("%q:\n got %s\nwant %s", c.text, got, c.want)
		}
	}
}
