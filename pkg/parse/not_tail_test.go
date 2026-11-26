//ff:func feature=parse type=reader control=sequence
//ff:what TestNotTail — the tail is read for the stem 아니 alone, each piece as one any predicate takes; the rest of the tail stays, and the tail it was made from is not changed
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNotTail(t *testing.T) {
	gate := &entry{name: "g"}
	from := tail{start: 2, gate: gate, noYo: true, pieces: []piece{{text: "란", names: "꾸밈:란(라는)", tail: rule.TailVerb}}}
	got := notTail(from)
	if got.only != notOnly || got.start != 2 || got.gate != gate || !got.noYo || len(got.pieces) != 1 || got.pieces[0].tail != rule.TailPred || got.pieces[0].names != "꾸밈:란(라는)" {
		t.Fatal(got)
	}
	if from.only != "" || from.pieces[0].tail != rule.TailVerb {
		t.Fatal("the tail it was made from", from)
	}
}
