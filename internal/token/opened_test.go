//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestOpened — a sign of the unit table right after a digit may be read as a unit of measure, and then the ways the signs stood open before it stay beside the ways after it; any other sign gives the ways after it as they are
package token

import (
	"reflect"
	"testing"
)

func TestOpened(t *testing.T) {
	ix, _ := newIndex(Options{})
	for _, c := range []struct {
		text    string
		at      int
		open    []string
		after   []string
		want    []string
		measure bool
	}{
		{"7\"", 1, []string{""}, []string{"\""}, []string{"\"", ""}, true},
		{"7'", 1, []string{"\""}, []string{"\"'"}, []string{"\"'", "\""}, true},
		{"가\"", 1, []string{""}, []string{"\""}, []string{"\""}, false},
		{"7)", 1, []string{""}, []string{"x"}, []string{"x"}, false},
		{"\"", 0, []string{""}, []string{"\""}, []string{"\""}, false},
	} {
		cut := newCutter(ix, c.text, false)
		cut.open = c.open
		got, measure := cut.opened(c.at, c.after)
		if !reflect.DeepEqual(got, c.want) || measure != c.measure {
			t.Errorf("%q: %q %v, want %q %v", c.text, got, measure, c.want, c.measure)
		}
	}
}
