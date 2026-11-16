//ff:func feature=token type=helper control=sequence
//ff:what TestDistinctPairing — each candidate of a sign of a pair once, in the order they first come (the order of the ways the signs before may stand open): no kind and no depth is put before another
package token

import (
	"reflect"
	"testing"
)

func TestDistinctPairing(t *testing.T) {
	got := distinct([]pairing{{shutName, 2}, {openName, 2}, {openName, 1}, {shutName, 1}, {openName, 1}, {shutName, 2}})
	want := []pairing{{shutName, 2}, {openName, 2}, {openName, 1}, {shutName, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	if got = distinct[pairing](nil); len(got) != 0 {
		t.Fatal("no candidates", got)
	}
}
