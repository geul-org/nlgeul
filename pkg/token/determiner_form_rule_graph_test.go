//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDeterminerFormRuleGraph — the graph of determinerFormRule evaluated by hand on rule.Around: 다른 written apart stands at the head of a sentence, after a nominal with another particle and after an adverb, and falls after a word ending in 와 or 과, and after 와 or 과 with a particle more (과는); 그런 after a word ending in 와 stands; written onto a word it falls, and written onto an opening sign it stands
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDeterminerFormRuleGraph(t *testing.T) {
	differ := rule.Near{Stem: "다른", Kind: "열쇠:관형사"}
	such := rule.Near{Stem: "그런", Kind: "열쇠:관형사"}
	glued := rule.Near{Stem: "다른", Kind: "열쇠:관형사", Glued: true}
	noun := rule.Near{Stem: "사람", Kind: "열쇠:체언", Last: "주격:이", Marks: "주격:이", Modifiable: true}
	with := func(last string) rule.Near {
		return rule.Near{Stem: "책", Kind: "열쇠:체언", Last: last, Marks: last, Nominal: true}
	}
	r := determinerFormRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"at the head", rule.Around{Self: differ, After: []rule.Near{noun}}, true},
		{"after the topic particle", rule.Around{Self: differ, Before: []rule.Near{with("보조사:은")}, After: []rule.Near{noun}}, true},
		{"after an adverb", rule.Around{Self: differ, Before: []rule.Near{{Stem: "또", Kind: "열쇠:부사"}}, After: []rule.Near{noun}}, true},
		{"after 과", rule.Around{Self: differ, Before: []rule.Near{with("부사격:과")}, After: []rule.Near{noun}}, false},
		{"after 와", rule.Around{Self: differ, Before: []rule.Near{with("부사격:와")}, After: []rule.Near{noun}}, false},
		{"after 과는", rule.Around{Self: differ, Before: []rule.Near{{Stem: "책", Kind: "열쇠:체언", Last: "보조사:는", Marks: "부사격:과+보조사:는", Nominal: true}}, After: []rule.Near{noun}}, false},
		{"그런 after 와", rule.Around{Self: such, Before: []rule.Near{with("부사격:와")}, After: []rule.Near{noun}}, true},
		{"onto a word", rule.Around{Self: glued, Before: []rule.Near{{Stem: "또", Kind: "열쇠:부사"}}, After: []rule.Near{noun}}, false},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}}, After: []rule.Near{noun}}, true},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
