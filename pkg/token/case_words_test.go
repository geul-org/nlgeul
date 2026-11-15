//ff:func feature=token type=helper control=sequence
//ff:what TestCaseWords — the words of rows in the order of the rows; no row, no word
package token

import (
	"strings"
	"testing"
)

func TestCaseWords(t *testing.T) {
	if got := caseWords([]caseException{{Word: "적소"}, {Word: "담소"}}); strings.Join(got, ",") != "적소,담소" {
		t.Fatal(got)
	}
	if got := caseWords(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
