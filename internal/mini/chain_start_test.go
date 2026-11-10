//ff:func feature=mini type=reader control=sequence
//ff:what TestChainStart — 만+으로+도 is read back to 만 with its particles and the case of 으로; no particle leaves the end; a fourth particle and a chain out of order (이+에) are not read
package mini

import "testing"

func TestChainStart(t *testing.T) {
	if got, kase, parts := chainStart([]rune("힘만으로도"), 5, 0, 5); got != 1 || kase != caseAdv || parts != "만+으로+도+" {
		t.Fatal(got, kase, parts)
	}
	if got, kase, parts := chainStart([]rune("매우"), 2, 0, 5); got != 2 || kase != 0 || parts != "" {
		t.Fatal(got, kase, parts)
	}
	if got, _, _ := chainStart([]rune("힘만으로도"), 5, 3, 5); got != 5 {
		t.Fatal(got)
	}
	if got, _, _ := chainStart([]rune("학생이에"), 4, 1, 5); got != 3 {
		t.Fatal("이+에 read as a chain", got)
	}
}
