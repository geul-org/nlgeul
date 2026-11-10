//ff:func feature=parse type=reader control=sequence
//ff:what TestStemMark — the copula stem 이 is a piece; 셨 after a stem is 시+었; a lone 셨 and other fused syllables stay as they are
package parse

import "testing"

func TestStemMark(t *testing.T) {
	if got := stemMark([]rune("방식이다"), 3, false, true); got != "방식+이" {
		t.Fatal(got)
	}
	if got := stemMark([]rune("가셨다"), 2, true, false); got != "가+시+었" {
		t.Fatal(got)
	}
	if got := stemMark([]rune("셨다"), 1, true, false); got != "셨" {
		t.Fatal(got)
	}
	if got := stemMark([]rune("지시했다"), 3, true, false); got != "지시+했" {
		t.Fatal(got)
	}
	if got := stemMark([]rune("먹었다"), 1, false, false); got != "먹" {
		t.Fatal(got)
	}
}
