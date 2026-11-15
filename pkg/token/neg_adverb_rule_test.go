//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestNegAdverbRule — the rule is named as the trace prints it, adds bonusSentence and waits for the word after the candidate; it speaks only of 안 or 못 read alone as an adverb of the word table (not of 안 the noun, not of 안 with a particle, not of another adverb); its claim stands when the word right after it is a verb or an adjective of the tables, and falls when that word is a noun, when no word follows, and when the word after it is a new word
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestNegAdverbRule(t *testing.T) {
	r := negAdverbRule()
	if r.name != ruleNegAdverb || r.bonus != bonusSentence || r.graph == nil || r.waits == nil {
		t.Fatalf("%+v", r)
	}
	adverb := rule.Near{Stem: "안", Kind: "열쇠:부사"}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{adverb, true},
		{rule.Near{Stem: "못", Kind: "열쇠:부사"}, true},
		{rule.Near{Stem: "안", Kind: "열쇠:체언"}, false},
		{rule.Near{Stem: "안", Kind: "열쇠:체언", Last: "부사격:에", Marks: "부사격:에"}, false},
		{rule.Near{Stem: "안", Kind: "열쇠:부사", Last: "보조사:도", Marks: "보조사:도"}, false},
		{rule.Near{Stem: "잘", Kind: "열쇠:부사"}, false},
	} {
		if got := r.about(c.near); got != c.want || r.waits(c.near) != c.want {
			t.Errorf("about(%+v) = %v", c.near, got)
		}
	}
	for _, c := range []struct {
		after []rule.Near
		want  bool
	}{
		{[]rule.Near{{Stem: "하", Kind: "열쇠:동사", Last: "맺음:어(합침)", Pred: true}}, true},
		{[]rule.Near{{Stem: "좋", Kind: "열쇠:형용사", Last: "맺음:다", Pred: true}}, true},
		{[]rule.Near{{Stem: "학교", Kind: "열쇠:체언"}}, false},
		{[]rule.Near{{Stem: "곰곰"}}, false},
		{nil, false},
	} {
		ctx := toulmin.NewContext()
		ctx.Set(rule.KeyAround, rule.Around{Self: adverb, After: c.after})
		if v, _ := evaluate(r.graph, ctx, false); (v > 0) != c.want {
			t.Errorf("after %+v: %v", c.after, v)
		}
	}
}
