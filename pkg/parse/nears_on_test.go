//ff:func feature=parse type=helper control=sequence
//ff:what TestNearsOn — the words after a seat come in the order of the line, the nearest first; no seat gives none
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNearsOn(t *testing.T) {
	got := nearsOn([]seat{{near: rule.Near{Stem: "a"}}, {near: rule.Near{Stem: "b", Pred: true}}})
	if len(got) != 2 || got[0].Stem != "a" || got[1] != (rule.Near{Stem: "b", Pred: true}) {
		t.Fatal(got)
	}
	if got := nearsOn(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
