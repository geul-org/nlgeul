//ff:func feature=token type=reader control=sequence
//ff:what TestPassedCandidates — the kinds that passed are read one by one under the table's kind in the order of kindOrder, each with its verdict; a closed-set word that passed as the kind of a set is read under the set's name, and also under the table's kind for each passed kind the word table lists for it, with the same verdict (the two names of one kind), and under the set's name alone when the table lists none of the passed kinds
package token

import (
	"reflect"
	"testing"
)

func TestPassedCandidates(t *testing.T) {
	vs := []float64{0.9, 0.8, 0.6, 0.5, 0.7, 0}
	got := passedCandidates(nil, kindNominal|kindAdverb, 0, vs)
	if want := []kindCandidate{{kindNominal, "열쇠:체언", 0.9}, {kindAdverb, "열쇠:부사", 0.5}}; !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	got = passedCandidates(nil, kindNominal|kindAdverb, kindNominal, vs)
	if !reflect.DeepEqual(got, []kindCandidate{{kindNominal, "열쇠:체언", 0.9}, {kindAdverb, "열쇠:부사", 0.5}}) {
		t.Fatal("a stem with no set's name is read under every passed kind whatever the table lists", got)
	}
	got = passedCandidates(&entry{name: "관형사:이"}, kindNominal|kindDeterminer, 0, vs)
	if want := []kindCandidate{{kindDeterminer, "관형사:이", 0.7}}; !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	got = passedCandidates(&entry{name: "관형사:이"}, kindNominal|kindDeterminer, kindAdverb, vs)
	if want := []kindCandidate{{kindDeterminer, "관형사:이", 0.7}}; !reflect.DeepEqual(got, want) {
		t.Fatal("a kind the table lists that did not pass is not read", got)
	}
	got = passedCandidates(&entry{name: "관형사:이"}, kindNominal|kindDeterminer, kindNominal, vs)
	if want := []kindCandidate{{kindDeterminer, "관형사:이", 0.7}, {kindNominal, "열쇠:체언", 0.9}}; !reflect.DeepEqual(got, want) {
		t.Fatal("a kind the table lists that passed is read under the table's name too", got)
	}
}
