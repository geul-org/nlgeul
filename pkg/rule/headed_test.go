//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestHeaded — applies (+1) when no word stands before the candidate (evidence "") or the stem's name of the word right before it has a kind in the list (열기 of 열기:1: an opening sign stands where a stem does, evidence the name); not (-1) when a word with another name stands there, a new word with no name, or a sign of another kind; the words further back are not read; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestHeaded(t *testing.T) {
	self := Near{Stem: "시", Kind: "열쇠:체언"}
	ctxOf := func(before ...Near) toulmin.Context {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, Around{Self: self, Before: before})
		return ctx
	}
	open := Near{Stem: "“", Kind: "열기:1"}
	cases := []struct {
		name string
		ctx  toulmin.Context
		not  bool
		want float64
		fact any
	}{
		{"head of the sentence", ctxOf(), false, 1, ""},
		{"no candidate", toulmin.NewContext(), false, 1, ""},
		{"right after an opening sign", ctxOf(open), false, 1, "열기:1"},
		{"after an opening sign after a word", ctxOf(open, Near{Stem: "열", Kind: "열쇠:동사", Last: "맺음:어", Pred: true, Marks: "맺음:어"}), false, 1, "열기:1"},
		{"after a noun", ctxOf(Near{Stem: "인구", Kind: "열쇠:체언"}), false, -1, "열쇠:체언"},
		{"after a new word", ctxOf(Near{Stem: "ᄒᆞᆫ"}), false, -1, ""},
		{"after a closing sign", ctxOf(Near{Stem: "”", Kind: "닫기:1"}), false, -1, "닫기:1"},
		{"after a comma", ctxOf(Near{Last: "쉼표", Marks: "쉼표"}), false, -1, ""},
		{"not, at the head", ctxOf(), true, -1, ""},
		{"not, after a noun", ctxOf(Near{Stem: "인구", Kind: "열쇠:체언"}), true, 1, "열쇠:체언"},
	}
	for _, c := range cases {
		specs := toulmin.Specs{&List{Name: "opening", Of: ListPrevKind, Items: []string{"열기"}}}
		if c.not {
			specs = append(specs, &Not{})
		}
		got, fact := Headed(c.ctx, specs)
		if got != c.want || fact != c.fact {
			t.Errorf("%s: Headed = %v, %v; want %v, %v", c.name, got, fact, c.want, c.fact)
		}
	}
}
