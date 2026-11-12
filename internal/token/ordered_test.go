//ff:func feature=token type=helper control=sequence
//ff:what TestOrdered — each reading once; those that open before those that close, each kind from the least depth up
package token

import (
	"reflect"
	"testing"
)

func TestOrdered(t *testing.T) {
	got := ordered([]pairing{{shutName, 2}, {openName, 2}, {openName, 1}, {shutName, 1}, {openName, 1}, {shutName, 2}})
	want := []pairing{{openName, 1}, {openName, 2}, {shutName, 1}, {shutName, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	if got = ordered(nil); len(got) != 0 {
		t.Fatal("no readings", got)
	}
}
