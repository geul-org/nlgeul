//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestApartParticleRuleGraph — the graph of apartParticleRule evaluated by hand on rule.Around: a particle written apart stands after a bare noun, at the head, after a closing sign, a dash, a run of Latin letters, a hyphened name and a code; it falls after a word ending in a particle, the copula or a closing, after an adverb or another word that ends in no nominal, and when it is written onto the word before, unless that word is a run of signs alone with no name (○○가), which stands for a name kept back; the vocative particle does not stand at the head (nothing before it to call) but stands after a bare noun or a closing sign
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestApartParticleRuleGraph(t *testing.T) {
	self := rule.Near{Last: "보조사:는", Marks: "보조사:는", Text: "는"}
	glued := self
	glued.Glued = true
	sign := func(kind string) []rule.Near { return []rule.Near{{Stem: "x", Kind: kind}} }
	vocative := rule.Near{Last: "호격:아", Marks: "호격:아", Text: "아"}
	r := apartParticleRule()
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"after a bare noun", rule.Around{Self: self, Before: []rule.Near{{Stem: "나무", Kind: "열쇠:체언", Nominal: true}}}, true},
		{"at the head", rule.Around{Self: self}, true},
		{"after a closing sign", rule.Around{Self: self, Before: sign("닫기:1")}, true},
		{"after a dash", rule.Around{Self: self, Before: []rule.Near{{Last: "줄표", Marks: "줄표"}}}, true},
		{"after a Latin run", rule.Around{Self: self, Before: sign("열:영숫자")}, true},
		{"after a hyphened name", rule.Around{Self: self, Before: sign("붙임:-")}, true},
		{"after a code", rule.Around{Self: self, Before: sign("코드:1")}, true},
		{"after a particle", rule.Around{Self: self, Before: []rule.Near{{Stem: "나무", Kind: "열쇠:체언", Last: "주격:가", Marks: "주격:가"}}}, false},
		{"after a closing", rule.Around{Self: self, Before: []rule.Near{{Stem: "자라", Kind: "열쇠:동사", Last: "맺음:다", Pred: true, Marks: "맺음:다"}}}, false},
		{"after the copula", rule.Around{Self: self, Before: []rule.Near{{Stem: "나무", Kind: "열쇠:체언", Last: "서술격:이", Marks: "서술격:이"}}}, false},
		{"after an adverb", rule.Around{Self: self, Before: []rule.Near{{Stem: "빨리", Kind: "열쇠:부사"}}}, false},
		{"after a full stop", rule.Around{Self: self, Before: sign("마침표")}, false},
		{"written onto a noun", rule.Around{Self: glued, Before: []rule.Near{{Stem: "나무", Kind: "열쇠:체언", Nominal: true}}}, false},
		{"written onto a run of signs", rule.Around{Self: glued, Before: []rule.Near{{Stem: "○○", Nominal: true, Signs: true}}}, true},
		{"written onto a run of signs that ends in a particle", rule.Around{Self: glued, Before: []rule.Near{{Stem: "○○", Last: "주격:가", Marks: "주격:가", Signs: true}}}, false},
		{"a vocative at the head", rule.Around{Self: vocative}, false},
		{"a vocative after a bare noun", rule.Around{Self: vocative, Before: []rule.Near{{Stem: "나무", Kind: "열쇠:체언", Nominal: true}}}, true},
		{"a vocative after a closing sign", rule.Around{Self: vocative, Before: sign("닫기:1")}, true},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
