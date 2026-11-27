//ff:func feature=parse type=helper control=sequence
//ff:what TestWithDefaults — a line of the table of kinds is added after the lines only for a letter they do not hold
package parse

import "testing"

func TestWithDefaults(t *testing.T) {
	list := []form{{'a', "x", "y", "a", "", "own"}}
	defaults := []form{{'a', "k", "k", "a", "", "kinds"}, {'b', "k", "k", "b", "", "kinds"}}
	got := withDefaults(list, defaults)
	if len(got) != 2 || got[0].origin != "own" || got[1].letter != 'b' {
		t.Fatal(got)
	}
}
