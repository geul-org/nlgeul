//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestBoundBeforeRuleGraph — the graph of boundBeforeRule evaluated by hand on rule.Around: a modifying form stands before 수, 것, 채 and 터 (the dependent nouns of the closed set and of the spacing words'), falls before a noun, at the sentence's end and before a dependent noun that stands after a nominal as well (때, 간, 만); a candidate that does not end in a modifying closing falls
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundBeforeRuleGraph(t *testing.T) {
	adn := rule.Near{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}
	next := func(stem, kind string) []rule.Near { return []rule.Near{{Stem: stem, Kind: kind}} }
	r := boundBeforeRule(boundNext, boundNominal)
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"before 수", rule.Around{Self: adn, After: next("수", "의존명사:수")}, true},
		{"before 수 read as a noun", rule.Around{Self: adn, After: next("수", "열쇠:체언")}, true},
		{"before 것", rule.Around{Self: adn, After: next("것", "의존명사:것")}, true},
		{"before 채", rule.Around{Self: adn, After: next("채", "의존명사:채")}, true},
		{"before 터", rule.Around{Self: adn, After: next("터", "열쇠:체언")}, true},
		{"before a noun", rule.Around{Self: adn, After: next("사과", "열쇠:체언")}, false},
		{"at the end", rule.Around{Self: adn}, false},
		{"before 때", rule.Around{Self: adn, After: next("때", "의존명사:때")}, false},
		{"before 간", rule.Around{Self: adn, After: next("간", "의존명사:간")}, false},
		{"before 만", rule.Around{Self: adn, After: next("만", "의존명사:만")}, false},
		{"a noun before 수", rule.Around{Self: rule.Near{Stem: "알", Kind: "열쇠:체언"}, After: next("수", "의존명사:수")}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
