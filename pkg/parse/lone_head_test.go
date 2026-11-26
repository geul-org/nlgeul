//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLoneHead — the heads of a path that are a path with no name after them (~/, ./, ../) are told at the front of a unit of signs, with their length; a slash alone, a head a slash or a dot goes on after, and other signs are none
package parse

import "testing"

func TestLoneHead(t *testing.T) {
	for in, want := range map[string]int{"~/": 2, "./": 2, "../": 3, "~/,": 2, "../)": 3, "/": 0, "//": 0, "~//": 0, "./.": 0, "../..": 0, "...": 0, "~": 0, "": 0, "->": 0} {
		if got := loneHead([]rune(in)); got != want {
			t.Errorf("%q: %d", in, got)
		}
	}
}
