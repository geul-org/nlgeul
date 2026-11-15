//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestKindListed — applies (+1, with the kind as evidence) when the part before the colon of the fact the list asks for is in the list (꾸밈 of 꾸밈:ㄹ as the candidate's last piece, 의존명사 of 의존명사:수 as its stem's name, 꾸밈 of 꾸밈:는 as the last piece of the word before it); not (-1) when the kind is another, when the fact is empty, or when the context holds no candidate; a name with no colon is its own kind; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestKindListed(t *testing.T) {
	ctxOf := func(a Around) toulmin.Context {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, a)
		return ctx
	}
	adn := Around{Self: Near{Stem: "알", Kind: "열쇠:동사", Last: "꾸밈:ㄹ", Pred: true, Marks: "꾸밈:ㄹ"}}
	bound := Around{Self: Near{Stem: "수", Kind: "의존명사:수"}, Before: []Near{{Stem: "하", Last: "꾸밈:는", Pred: true}}}
	cases := []struct {
		name  string
		ctx   toulmin.Context
		of    int
		items []string
		not   bool
		want  float64
		fact  any
	}{
		{"self last", ctxOf(adn), ListSelfLast, []string{"꾸밈"}, false, 1, "꾸밈"},
		{"self stem's name", ctxOf(bound), ListSelfKind, []string{"의존명사"}, false, 1, "의존명사"},
		{"previous last", ctxOf(bound), ListPrevLast, []string{"꾸밈", "관형격"}, false, 1, "꾸밈"},
		{"another kind", ctxOf(adn), ListSelfLast, []string{"맺음"}, false, -1, "꾸밈"},
		{"the whole name is no kind", ctxOf(adn), ListSelfLast, []string{"꾸밈:ㄹ"}, false, -1, "꾸밈"},
		{"no colon", ctxOf(Around{Self: Near{Last: "쉼표"}}), ListSelfLast, []string{"쉼표"}, false, 1, "쉼표"},
		{"empty fact", ctxOf(adn), ListPrevLast, []string{"꾸밈"}, false, -1, ""},
		{"no candidate", toulmin.NewContext(), ListSelfLast, []string{"꾸밈"}, false, -1, ""},
		{"not, held", ctxOf(adn), ListSelfLast, []string{"꾸밈"}, true, -1, "꾸밈"},
		{"not, missing", ctxOf(adn), ListSelfLast, []string{"맺음"}, true, 1, "꾸밈"},
	}
	for _, c := range cases {
		specs := toulmin.Specs{&List{Name: "kinds", Of: c.of, Items: c.items}}
		if c.not {
			specs = append(specs, &Not{})
		}
		got, fact := KindListed(c.ctx, specs)
		if got != c.want || fact != c.fact {
			t.Errorf("%s: KindListed = %v, %v; want %v, %v", c.name, got, fact, c.want, c.fact)
		}
	}
}
