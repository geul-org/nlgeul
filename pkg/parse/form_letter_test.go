//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestFormLetter — a letter of the tables of special letters that is no sign and no letter of a run is read as a letter of its own (①, ½, ㊀, ㈠); a sign of the tables (ⓐ, ㈎, ？), a letter of a run (Ⅲ, Ａ, １) and a letter the tables lack (가, 1, ○) are not
package parse

import "testing"

func TestFormLetter(t *testing.T) {
	for r, want := range map[rune]bool{'①': true, '½': true, '㊀': true, '㈠': true, 'ⓐ': false, '㈎': false, '？': false, 'Ⅲ': false, 'Ａ': false, '１': false, '가': false, '1': false, '○': false} {
		if got := formLetter(r); got != want {
			t.Errorf("formLetter(%c) = %v", r, got)
		}
	}
}
