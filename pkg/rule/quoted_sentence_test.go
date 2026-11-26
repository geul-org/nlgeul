//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestQuotedSentence — applies (+1, with the name of the closing as evidence) when the word right before the candidate is a closing mark and the word before that, or before the sign that ends a sentence between them, ends in a closing; not (-1) when that word ends in a nominal, when the word before is no closing mark, when the sign that ends a sentence is the only word before the mark, or when there is no word or no candidate; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestQuotedSentence(t *testing.T) {
	l := &List{Name: "stops", Of: ListPrevLast, Items: []string{"마침표", "물음표"}}
	shut := Near{Kind: "닫기:1", Last: "닫기:1"}
	ctxOf := func(before ...Near) toulmin.Context {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: Near{Marks: "꾸밈:는", Last: "꾸밈:는"}, Before: before})
		return ctx
	}
	closing := Near{Stem: "오", Kind: "열쇠:동사", Last: "맺음:ㄴ다", Marks: "맺음:ㄴ다"}
	noun := Near{Stem: "이름", Kind: "열쇠:체언"}
	cases := []struct {
		name string
		ctx  toulmin.Context
		not  bool
		want float64
		fact any
	}{
		{"closing", ctxOf(shut, closing), false, 1, "맺음:ㄴ다"},
		{"stop between", ctxOf(shut, Near{Last: "물음표"}, closing), false, 1, "맺음:ㄴ다"},
		{"noun", ctxOf(shut, noun), false, -1, ""},
		{"no mark before", ctxOf(closing, closing), false, -1, nil},
		{"stop alone", ctxOf(shut, Near{Last: "마침표"}), false, -1, ""},
		{"no word before the mark", ctxOf(shut), false, -1, ""},
		{"no word", ctxOf(), false, -1, nil},
		{"no candidate", toulmin.NewContext(), false, -1, nil},
		{"not, held", ctxOf(shut, closing), true, -1, "맺음:ㄴ다"},
		{"not, missing", ctxOf(shut, noun), true, 1, ""},
	}
	for _, c := range cases {
		specs := toulmin.Specs{l}
		if c.not {
			specs = toulmin.Specs{l, &Not{}}
		}
		got, fact := QuotedSentence(c.ctx, specs)
		if got != c.want || fact != c.fact {
			t.Errorf("%s: %v %q, want %v %q", c.name, got, fact, c.want, c.fact)
		}
	}
}
