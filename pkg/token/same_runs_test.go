//ff:func feature=token type=helper control=sequence
//ff:what TestSameRuns — signs are cut where the sign changes, each stretch as where it starts and ends; one sign is one stretch; nothing gives nothing
package token

import (
	"reflect"
	"testing"
)

func TestSameRuns(t *testing.T) {
	if got := sameRuns([]rune("?!!..-")); !reflect.DeepEqual(got, [][2]int{{0, 1}, {1, 3}, {3, 5}, {5, 6}}) {
		t.Fatal(got)
	}
	if got := sameRuns([]rune("...")); !reflect.DeepEqual(got, [][2]int{{0, 3}}) {
		t.Fatal(got)
	}
	if got := sameRuns(nil); got != nil {
		t.Fatal(got)
	}
}
