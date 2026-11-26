//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCalls — a comma (the fullwidth one too) and an exclamation mark at the head of a unit's last signs are written after one who is called; a full stop, a question mark, other signs and no sign at all are not
package parse

import "testing"

func TestCalls(t *testing.T) {
	cases := []struct {
		signs string
		want  bool
	}{
		{",", true}, {"，", true}, {"!", true}, {"!\"", true}, {",\"", true}, {"!!", true},
		{"", false}, {".", false}, {"?", false}, {"\"", false}, {"\",", false}, {"…", false}, {")!", false},
	}
	for _, c := range cases {
		if got := calls(c.signs); got != c.want {
			t.Errorf("calls(%q) = %v, want %v", c.signs, got, c.want)
		}
	}
}
