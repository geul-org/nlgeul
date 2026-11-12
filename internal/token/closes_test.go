//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCloses — a question mark, an exclamation mark and a single full stop at the head of a unit's last signs end a sentence; dots in a row, other signs and no sign at all do not
package token

import "testing"

func TestCloses(t *testing.T) {
	cases := []struct {
		signs string
		want  bool
	}{
		{"?", true}, {"!", true}, {".", true}, {"?!", true}, {".\"", true}, {"!)", true},
		{"", false}, {"..", false}, {"...", false}, {",", false}, {")", false}, {")?", false}, {"…", false},
	}
	for _, c := range cases {
		if got := closes(c.signs); got != c.want {
			t.Errorf("closes(%q) = %v, want %v", c.signs, got, c.want)
		}
	}
}
