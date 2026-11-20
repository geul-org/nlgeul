//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestEnclosedNumber — the first and last letter of each of the four ranges are numbers; the letters just outside them, circled Hangul and Latin letters, and a digit are not
package parse

import "testing"

func TestEnclosedNumber(t *testing.T) {
	for _, r := range []rune{'①', '⑳', '⑴', '⒇', '㉑', '㉟', '㊱', '㊿', '⑦'} {
		if !enclosedNumber(r) {
			t.Errorf("%c is a number", r)
		}
	}
	for _, r := range []rune{'⓪', '⒈', '㉠', '㉮', 'ⓐ', '⒜', '❶', '⓫', '㊿' + 1, '㉟' + 1, '1', '가'} {
		if enclosedNumber(r) {
			t.Errorf("%c is no number", r)
		}
	}
}
