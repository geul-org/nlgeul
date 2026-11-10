//ff:func feature=token type=reader control=sequence
//ff:what TestClosingDots — a full stop or dots right after the core are read as a line; a unit that ends otherwise has none
package token

import "testing"

func TestClosingDots(t *testing.T) {
	sc := scanOf("올렸다")
	if got := sc.closingDots("올렸다."); len(got) != 1 || got[0] != ".[마침표:.]" {
		t.Fatal(got)
	}
	if got := scanOf("글쎄").closingDots("\"글쎄...\""); len(got) != 1 || got[0] != "...[줄임표:.]" {
		t.Fatal(got)
	}
	if got := sc.closingDots("올렸다?"); got != nil {
		t.Fatal(got)
	}
	if got := sc.closingDots("올렸다"); got != nil {
		t.Fatal(got)
	}
}
