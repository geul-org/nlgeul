//ff:func feature=token type=helper control=sequence
//ff:what TestNearsBack — the words before a seat come nearest first: the seats given from the last back, then the words judged earlier as they are kept (already nearest first)
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNearsBack(t *testing.T) {
	seats := []seat{{near: rule.Near{Stem: "a"}}, {near: rule.Near{Stem: "b"}}}
	behind := []rule.Near{{Stem: "c"}, {Stem: "d"}}
	got := nearsBack(seats, behind)
	if len(got) != 4 || got[0].Stem != "b" || got[1].Stem != "a" || got[2].Stem != "c" || got[3].Stem != "d" {
		t.Fatal(got)
	}
	if got := nearsBack(nil, behind); len(got) != 2 || got[0].Stem != "c" {
		t.Fatal(got)
	}
	if got := nearsBack(nil, nil); len(got) != 0 {
		t.Fatal(got)
	}
}
