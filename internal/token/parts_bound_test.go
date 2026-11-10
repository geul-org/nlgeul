//ff:func feature=token type=helper control=sequence
//ff:what TestPartsBound — the words between slashes, or between hyphens; none for no sign, both kinds, or a sign at an edge or doubled
package token

import "testing"

func TestPartsBound(t *testing.T) {
	if got := partsBound("가/나/다"); len(got) != 3 || got[2] != "다" {
		t.Fatal(got)
	}
	if got := partsBound("가-나"); len(got) != 2 {
		t.Fatal(got)
	}
	if partsBound("가나") != nil || partsBound("가/나-다") != nil || partsBound("/가") != nil || partsBound("가//나") != nil {
		t.Fatal("not bound words")
	}
}
