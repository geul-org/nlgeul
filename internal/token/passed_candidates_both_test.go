//ff:func feature=token type=reader control=sequence
//ff:what TestPassedCandidatesBoth — a closed-set word is read under the set's name and under the table's kind for each passed kind the table lists, the set's names first (in the order of the name) and the table's kinds after them in kindOrder, each with its kind's verdict; a set's kind the table does not list is named by the set alone
package token

import (
	"reflect"
	"testing"
)

func TestPassedCandidatesBoth(t *testing.T) {
	vs := []float64{0.9, 0.8, 0.6, 0.5, 0.7, 0}
	got := passedCandidates(&entry{name: "의존명사:수"}, kindNominal, kindNominal|kindVerb, vs)
	if want := []kindCandidate{{kindNominal, "의존명사:수", 0.9}, {kindNominal, "열쇠:체언", 0.9}}; !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	got = passedCandidates(&entry{name: "보조 용언:있"}, kindVerb|kindAdj, kindVerb|kindAdj, vs)
	want := []kindCandidate{{kindVerb, "보조 용언:있", 0.8}, {kindVerb, "열쇠:동사", 0.8}, {kindAdj, "열쇠:형용사", 0.6}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	got = passedCandidates(&entry{name: "보조 용언:있"}, kindVerb|kindAdj, kindAdj, vs)
	if want := []kindCandidate{{kindVerb, "보조 용언:있", 0.8}, {kindAdj, "열쇠:형용사", 0.6}}; !reflect.DeepEqual(got, want) {
		t.Fatal("the table lists the adjective alone: the verb is named by the set alone", got)
	}
	got = passedCandidates(&entry{name: "관형사:한|의존명사:한"}, kindNominal|kindDeterminer, kindNominal|kindDeterminer, vs)
	want = []kindCandidate{{kindDeterminer, "관형사:한", 0.7}, {kindNominal, "의존명사:한", 0.9}, {kindNominal, "열쇠:체언", 0.9}, {kindDeterminer, "열쇠:관형사", 0.7}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
}
