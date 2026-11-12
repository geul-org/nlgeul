//ff:func feature=token type=helper control=sequence
//ff:what TestKindCandidatesOf — each candidate of a sign becomes a candidate of kind 0 with its own name and verdict, in order; no candidates, none
package token

import (
	"reflect"
	"testing"
)

func TestKindCandidatesOf(t *testing.T) {
	got := kindCandidatesOf([]markCandidate{{"파일명", 1}, {"도메인", 0.5}})
	want := []kindCandidate{{0, "파일명", 1}, {0, "도메인", 0.5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatal(got)
	}
	if got := kindCandidatesOf(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
