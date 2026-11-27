//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestEnclosedNumber — the letters the tables of special letters give the shape 원 or 괄호 and the content 수 are numbers (①, ⑳, ⑴, ⒇, ㉑, ㊿, and ⓪ and ➀ too); a number of another shape, circled Hangul and Latin letters, and a digit are not
package parse

import "testing"

func TestEnclosedNumber(t *testing.T) {
	for _, r := range []rune{'①', '⑳', '⑴', '⒇', '㉑', '㉟', '㊱', '㊿', '⑦', '⓪', '➀'} {
		if !enclosedNumber(r) {
			t.Errorf("%c is a number", r)
		}
	}
	for _, r := range []rune{'⒈', '㉠', '㉮', 'ⓐ', '⒜', '❶', '⓫', '㊿' + 1, '㉟' + 1, '1', '가'} {
		if enclosedNumber(r) {
			t.Errorf("%c is no number", r)
		}
	}
}
