//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNumberNameOf — a number is named by its letters: a digit among them (a comma, a point, a sign) 수:십진; only Hangul numerals (사십, 열두, 수십) 수:한글십진; Hangul that is no numeral (몇, 가) the plain 수; a circled number 원:수; a parenthesized number 괄호:수; one Latin letter 특수:로마자; only Chinese characters 수:한자십진; only Roman numeral letters 수:로마십진; a circled Latin letter, a mix of scripts and no letters the plain 수
package parse

import "testing"

func TestNumberNameOf(t *testing.T) {
	for text, want := range map[string]string{
		"1000": "수:십진", "2,473": "수:십진", "78,392.03": "수:십진", "-2": "수:십진", ".238": "수:십진", "+60": "수:십진", "제1": "수:십진",
		"한": "수:한글십진", "삼십": "수:한글십진", "몇": "수", "가": "수", "여러": "수", "사십": "수:한글십진", "수십": "수:한글십진", "열두": "수:한글십진", "스물한": "수:한글십진", "가나": "수", "수": "수",
		"三十": "수:한자십진", "〇": "수:한자십진",
		"Ⅳ": "수:로마십진", "ⅻ": "수:로마십진",
		"①": "원:수", "㊿": "원:수", "⑴": "괄호:수", "n": "특수:로마자", "K": "특수:로마자", "ⓛ": "수", "삼n": "수", "三삼": "수", "": "수", "제①": "수",
	} {
		if got := numberNameOf(text); got != want {
			t.Errorf("%q: %s, want %s", text, got, want)
		}
	}
}
