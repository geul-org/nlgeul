//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestHangulNumeral — a run of the numerals written in Hangul is one (하나, 사십, 열두, 스물한, 두세, 육칠만, 이만, 수십, 수백만); other Hangul (몇, 여러, 가, 수), a numeral with a letter that is none, an empty text and a 수 not before a greater number are not
package parse

import "testing"

func TestHangulNumeral(t *testing.T) {
	for text, want := range map[string]bool{
		"하나": true, "사십": true, "열두": true, "스물한": true, "두세": true, "육칠만": true, "이만": true, "수십": true, "수백만": true, "삼십칠": true, "한": true,
		"몇": false, "여러": false, "가": false, "수": false, "수일": false, "사십가": false, "": false, "나": false, "다섯째": false,
	} {
		if got := hangulNumeral(text); got != want {
			t.Errorf("%q: %v, want %v", text, got, want)
		}
	}
}
