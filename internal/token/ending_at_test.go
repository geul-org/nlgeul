//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestEndingAt — the states that end at a place, and none where none ends
package token

import "testing"

func TestEndingAt(t *testing.T) {
	r, _ := windowSteps(t, 5, 13)
	got := endingAt(r.byEnd(), 4)
	if len(got) == 0 {
		t.Fatal("none ends at 4")
	}
	for s := range got {
		if s.last.end != 4 {
			t.Fatal(s.last.end)
		}
	}
	if len(endingAt(r.byEnd(), 99)) != 0 {
		t.Fatal("none")
	}
}
