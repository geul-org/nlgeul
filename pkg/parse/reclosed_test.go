//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReclosed — of the closings read at the end of 먹었다거나, 거나 is read once more after the final closing 다 (먹+었+다+거나); the closings read at the end of 먹었고 give no such tail
package parse

import (
	"slices"
	"testing"
)

func TestReclosed(t *testing.T) {
	sc := scanOf("먹었다거나")
	pieces, from, starts := sc.readGroup(classClosing, 5, -1, false)
	got := sc.reclosed(pieces, from, starts, 5)
	if !slices.ContainsFunc(got, func(l tail) bool {
		n := len(l.pieces)
		return n >= 2 && l.pieces[n-2].names == "맺음:다" && l.pieces[n-1].names == "맺음:거나"
	}) {
		t.Fatalf("no 다+거나 %+v", got)
	}
	sc = scanOf("먹었고")
	pieces, from, starts = sc.readGroup(classClosing, 3, -1, false)
	for _, l := range sc.reclosed(pieces, from, starts, 3) {
		t.Errorf("%+v", l)
	}
}
