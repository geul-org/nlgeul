//ff:func feature=parse type=helper control=sequence
//ff:what TestWaysOf — each candidate of a sign of a pair with the way after it, in their order; none gives none
package parse

import (
	"reflect"
	"testing"
)

func TestWaysOf(t *testing.T) {
	got := waysOf([]pairing{{openName, 1, 0.5}, {shutName, 0, 1}}, []string{"\"", ""})
	want := []pairWay{{pairing{openName, 1, 0.5}, "\""}, {pairing{shutName, 0, 1}, ""}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	if got = waysOf(nil, nil); len(got) != 0 {
		t.Fatal(got)
	}
}
