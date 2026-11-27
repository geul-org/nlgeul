//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberScript — a digit is 십진, a Hangul syllable 한글십진, a Chinese character (and 〇) 한자십진, a Roman numeral letter 로마십진; any other letter (circled number, Latin letter, sign, a lone Hangul letter) none
package parse

import "testing"

func TestNumberScript(t *testing.T) {
	for r, want := range map[rune]string{
		'0': "십진", '9': "십진", '가': "한글십진", '힣': "한글십진", '一': "한자십진", '龥': "한자십진", '㐀': "한자십진", '〇': "한자십진",
		'Ⅰ': "로마십진", 'ⅿ': "로마십진", '①': "", 'n': "", '-': "", 'ㄱ': "",
	} {
		if got := numberScript(r); got != want {
			t.Errorf("%q: %q, want %q", r, got, want)
		}
	}
}
