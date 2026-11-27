//ff:func feature=parse type=helper control=sequence
//ff:what TestHasNumeralText — a list of numbers has a text when one of them is written so
package parse

import "testing"

func TestHasNumeralText(t *testing.T) {
	words := []numeral{{"삼", -1, numName}, {"사", -1, numName}}
	if !hasNumeralText(words, "사") || hasNumeralText(words, "오") || hasNumeralText(nil, "삼") {
		t.Fatal("has")
	}
}
