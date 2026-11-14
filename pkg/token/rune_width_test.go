//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestRuneWidth — a letter of Hangul, kana, a CJK ideograph, a full width form or an emoji takes two cells, a combining mark or a zero width letter none, a Latin letter, a digit, a box drawing letter and a circled number one
package token

import "testing"

func TestRuneWidth(t *testing.T) {
	for r, want := range map[rune]int{
		'가': 2, 'ㄴ': 2, 'あ': 2, '漢': 2, 'Ａ': 2, '😀': 2, '힣': 2,
		'́': 0, '​': 0, '️': 0,
		'a': 1, '7': 1, '─': 1, '│': 1, '┐': 1, '①': 1, 'λ': 1, ' ': 1,
	} {
		if got := runeWidth(r); got != want {
			t.Fatalf("%q: %d, want %d", r, got, want)
		}
	}
}
