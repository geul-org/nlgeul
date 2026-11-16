//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDeterminerAloneRuleGraph — the graph of determinerAloneRule evaluated by hand on rule.Around: the determiner 이 written apart stands before a noun, a determiner, a modifying form and an opening sign, and falls before an adverb, before a predicate, before a sign other than an opening one and at the end of the sentence; written onto a word it falls (before a noun too), and written onto an opening sign it stands
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestDeterminerAloneRuleGraph(t *testing.T) {
	apart := rule.Near{Stem: "이", Kind: "관형사:이"}
	glued := rule.Near{Stem: "이", Kind: "관형사:이", Glued: true}
	noun := rule.Near{Stem: "책", Kind: "열쇠:체언", Last: "목적격:을", Marks: "목적격:을", Modifiable: true}
	open := rule.Near{Stem: "“", Kind: "열기:1"}
	r := determinerAloneRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before a noun", rule.Around{Self: apart, After: []rule.Near{noun}}, true},
		{"before a determiner", rule.Around{Self: apart, After: []rule.Near{{Stem: "한", Kind: "관형사:한", Modifiable: true}}}, true},
		{"before a modifying form", rule.Around{Self: apart, After: []rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "꾸밈:은", Pred: true, Modifiable: true}}}, true},
		{"before an opening sign", rule.Around{Self: apart, After: []rule.Near{open, noun}}, true},
		{"before an adverb", rule.Around{Self: apart, After: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}, noun}}, false},
		{"before a predicate", rule.Around{Self: apart, After: []rule.Near{{Stem: "읽", Kind: "열쇠:동사", Last: "맺음:다", Pred: true}}}, false},
		{"before a closing sign", rule.Around{Self: apart, After: []rule.Near{{Stem: "”", Kind: "닫기:1"}}}, false},
		{"at the end", rule.Around{Self: apart}, false},
		{"onto a word", rule.Around{Self: glued, Before: []rule.Near{{Stem: "하나", Kind: "열쇠:체언", Last: "보조사:도"}}, After: []rule.Near{noun}}, false},
		{"onto an opening sign", rule.Around{Self: glued, Before: []rule.Near{open}, After: []rule.Near{noun}}, true},
		{"onto an opening sign before an adverb", rule.Around{Self: glued, Before: []rule.Near{open}, After: []rule.Near{{Stem: "다시", Kind: "열쇠:부사"}}}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
