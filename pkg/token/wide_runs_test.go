//ff:func feature=token type=reader control=sequence
//ff:what TestWideRuns — only the runs of digits and a unit are listed, from their start to their end; a unit alone, digits alone and another run are not
package token

import "testing"

func TestWideRuns(t *testing.T) {
	s := []rune("7m kg 45 v2 64GB")
	got := wideRuns(s, runsOf(s))
	if len(got) != 2 || got[0] != 2 || got[12] != 16 {
		t.Fatal(got)
	}
}
