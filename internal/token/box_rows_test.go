//ff:func feature=token type=helper control=sequence
//ff:what TestBoxRows — a row for each candidate: its letters padded to the widest, 성립 and the verdict, 선호 and the preference set right; with a head the first ones only and a last row that counts the others (head as many as there are, or 0, leaves none out)
package token

import (
	"slices"
	"testing"
)

func TestBoxRows(t *testing.T) {
	n := boxNode(0, 3, "아버지", "아버지", "아버", "아", "a")
	n.Contest.Branches[3].Pref = -113.5
	all := []string{
		"아버지  성립 1.00 선호   -3.25",
		"아버    성립 0.75 선호   -6.50",
		"아      성립 0.50 선호   -9.75",
		"a       성립 0.25 선호 -113.50",
	}
	if got := boxRows(n, 0); !slices.Equal(got, all) {
		t.Fatalf("%q", got)
	}
	if got := boxRows(n, 4); !slices.Equal(got, all) {
		t.Fatalf("%q", got)
	}
	want := []string{
		"아버지  성립 1.00 선호 -3.25",
		"아버    성립 0.75 선호 -6.50",
		"(2개 생략됨)",
	}
	if got := boxRows(n, 2); !slices.Equal(got, want) {
		t.Fatalf("%q", got)
	}
	if got := boxRows(n, 9); !slices.Equal(got, all) {
		t.Fatalf("%q", got)
	}
}
