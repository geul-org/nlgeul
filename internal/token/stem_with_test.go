//ff:func feature=token type=helper control=sequence
//ff:what TestStemWith — each candidate of a stem is given the stem's letters and the letters it is written with, in order; none for no candidates
package token

import (
	"reflect"
	"testing"
)

func TestStemWith(t *testing.T) {
	candidates := []kindCandidate{{kindNominal, "열쇠:체언", 1}, {kindAdj, "열쇠:형용사", 0.5}}
	got := stemWith("역할", "역활", candidates)
	want := []stemCandidate{{candidates[0], "역할", "역활"}, {candidates[1], "역할", "역활"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	if got := stemWith("곰", "", nil); len(got) != 0 {
		t.Fatal(got)
	}
}
