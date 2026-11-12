//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddPairs — each sign of a pair gets a graph for each of its candidates: an opening bracket one, a closing bracket one, a quotation mark written the same way at both ends two (열기 first); 22 signs give 24 graphs
package token

import (
	"reflect"
	"testing"
)

func TestAddPairs(t *testing.T) {
	ix := &index{}
	ix.addPairs()
	if len(ix.pairs) != len(pairOrder) || ix.graphs != 24 {
		t.Fatal(len(ix.pairs), ix.graphs)
	}
	for sign, want := range map[rune][]string{'(': {"열기:("}, ')': {"닫기:)"}, '"': {"열기:\"", "닫기:\""}, '\'': {"열기:'", "닫기:'"}, '》': {"닫기:》"}} {
		if got := entryNames(ix.pairs[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
}
