//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosedBefore — 거나 after 먹었다 is read once more as an ending after the final closing 다: the tail of 었+다 (or 다 alone) with 거나 as one more piece of a predicate tail, keeping the stem's closing and taking 거나's 요; nothing for an ending that is not in closingChains (고), at the head of the word, or after a closing it is not read after (먹고거나)
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosedBefore(t *testing.T) {
	sc := scanOf("먹었다거나")
	e := sc.ix.named('나', "맺음:거나")
	pc := piece{text: "거나", names: "맺음:거나", verdict: 1, tail: rule.TailVerb}
	got := sc.closedBefore(pc, e, 3, 5)
	i := slices.IndexFunc(got, func(l tail) bool { return len(l.pieces) == 3 && l.start == 1 })
	if i < 0 {
		t.Fatalf("no 었+다+거나 %+v", got)
	}
	l := got[i]
	if l.pieces[1].names != "맺음:다" || l.pieces[2].names != "맺음:거나" || l.pieces[2].tail != rule.TailPred || l.close == nil || l.close.name != "맺음:다" || l.noYo != e.noYo {
		t.Fatalf("%+v", l)
	}
	if pc.tail != rule.TailVerb {
		t.Fatal("the closing's own piece is left as it was")
	}
	for _, c := range []struct {
		text  string
		name  string
		start int
	}{
		{"먹었다고", "맺음:고", 3}, {"다거나", "맺음:거나", 1}, {"먹고거나", "맺음:거나", 2},
	} {
		sc := scanOf(c.text)
		pc := piece{text: string([]rune(c.text)[c.start:]), names: c.name, verdict: 1}
		if got := sc.closedBefore(pc, sc.ix.named([]rune(c.text)[len([]rune(c.text))-1], c.name), c.start, len([]rune(c.text))); len(got) != 0 {
			t.Errorf("%s: %+v", c.text, got)
		}
	}
}
