//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberLetter — the letters the tables of special letters give the content 수 in a shape of their own are numbers of one letter (①, ⑳, ⑴, ⒇, ㉑, ㊿, ⓪, ➀, ⒈, ❶, ⓫, ½, ²); circled Hangul and Latin letters and a circled Chinese numeral (another content), a full-width digit and a Roman numeral letter (letters of a run), a digit and a Hangul syllable are not
package parse

import "testing"

func TestNumberLetter(t *testing.T) {
	for _, r := range []rune{'①', '⑳', '⑴', '⒇', '㉑', '㉟', '㊱', '㊿', '⑦', '⓪', '➀', '⒈', '❶', '⓫', '½', '²'} {
		if !numberLetter(r) {
			t.Errorf("%c is a number", r)
		}
	}
	for _, r := range []rune{'㉠', '㉮', 'ⓐ', '⒜', '㊀', '１', 'Ⅲ', '1', '가'} {
		if numberLetter(r) {
			t.Errorf("%c is no number", r)
		}
	}
}
