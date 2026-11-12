//ff:func feature=token type=reader control=sequence
//ff:what TestPassedReads — the kinds that passed are read one by one under the table's kind in the order of kindOrder, each with its verdict; a closed-set word that passed as the kind of a set is read under the set's name only, with the verdict of that kind
package token

import (
	"reflect"
	"testing"
)

func TestPassedReads(t *testing.T) {
	vs := []float64{0.9, 0, 0, 0.5, 0.7, 0}
	got := passedReads(nil, kindNominal|kindAdverb, vs)
	if want := []kindRead{{kindNominal, "열쇠:체언", 0.9}, {kindAdverb, "열쇠:부사", 0.5}}; !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	got = passedReads(&entry{name: "관형사:이"}, kindNominal|kindDeterminer, vs)
	if want := []kindRead{{kindDeterminer, "관형사:이", 0.7}}; !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
}
