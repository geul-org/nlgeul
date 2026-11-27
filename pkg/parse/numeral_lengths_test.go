//ff:func feature=parse type=helper control=sequence
//ff:what TestNumeralLengths — the lengths of the numerals a text begins with (삼십 begins with 삼; 열 and 열두 are not both: only 열 of the table), and none when no run ends before the text
package parse

import "testing"

func TestNumeralLengths(t *testing.T) {
	if got := numeralLengths("삼십", true); len(got) != 1 || got[0] != len("삼") {
		t.Fatal(got)
	}
	if len(numeralLengths("삼십", false)) != 0 || len(numeralLengths("몇", true)) != 0 {
		t.Fatal("none")
	}
}
