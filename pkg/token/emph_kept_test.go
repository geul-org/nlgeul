//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEmphKept — for 를 of emphasis only the chains that end in 에, 으로 or 로 are kept, in their order; for any other particle every chain is kept as it is
package token

import (
	"slices"
	"testing"
)

func TestEmphKept(t *testing.T) {
	ends := []string{"부사격:에", "부사격:서", "부사격:으로", "부사격:에서", "부사격:로"}
	var tails []tail
	for i, name := range ends {
		tails = append(tails, tail{start: i, pieces: []piece{{text: "x", names: name}}})
	}
	got := emphKept(&entry{name: emphObjName}, tails)
	var names []string
	for _, l := range got {
		names = append(names, l.pieces[0].names)
	}
	if !slices.Equal(names, []string{"부사격:에", "부사격:으로", "부사격:로"}) || got[0].start != 0 || got[1].start != 2 {
		t.Fatal(names, got)
	}
	if got := emphKept(&entry{name: emphName}, tails); len(got) != len(tails) {
		t.Fatal("the emphatic 가", got)
	}
	if got := emphKept(&entry{name: emphObjName}, nil); got != nil {
		t.Fatal("no chains", got)
	}
}
