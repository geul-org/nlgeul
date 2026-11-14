//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestCellWidth — a text takes the sum of its letters' cells: Hangul two each, ASCII and the box drawing letters one, a combining mark none, the empty text none
package token

import "testing"

func TestCellWidth(t *testing.T) {
	for s, want := range map[string]int{
		"": 0, "abc": 3, "아버지": 6, "a가b": 4, "성립 1.00 선호 ": 15, "┌─┐": 3, "é": 1, "(ㄴ)다": 6,
	} {
		if got := cellWidth(s); got != want {
			t.Fatalf("%q: %d, want %d", s, got, want)
		}
	}
}
