//ff:func feature=parse type=helper control=sequence
//ff:what TestNamedEntries — the entries with a name are kept in their order and one with no name is left out; no entries give none
package parse

import (
	"reflect"
	"testing"
)

func TestNamedEntries(t *testing.T) {
	a, b, c := &entry{name: "줄표"}, &entry{}, &entry{name: "빼기"}
	if got := namedEntries([]*entry{a, b, c}); !reflect.DeepEqual(got, []*entry{a, c}) {
		t.Fatal(got)
	}
	if got := namedEntries(nil); got != nil {
		t.Fatal(got)
	}
}
