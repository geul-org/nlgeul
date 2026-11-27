//ff:func feature=parse type=reader control=sequence
//ff:what TestHanNumerals — the number written in Han characters that ends at a place is the whole run of Han numerals there, named 수:한자십진 and valued letter by letter; none where no Han numeral ends, or where another Han character stands right before the run (第一, 張三)
package parse

import "testing"

func TestHanNumerals(t *testing.T) {
	got := scanOf("二三일").hanNumerals(2)
	if len(got) != 1 || got[0].text != "二三" || got[0].name != numName || got[0].pref >= 0 {
		t.Fatal(got)
	}
	one := scanOf("十원").hanNumerals(1)
	if len(one) != 1 || one[0].text != "十" || one[0].pref <= got[0].pref {
		t.Fatal(one, got)
	}
	if got = scanOf("곰二일").hanNumerals(2); len(got) != 1 || got[0].text != "二" {
		t.Fatal(got)
	}
	if got = scanOf("第一회").hanNumerals(2); got != nil {
		t.Fatal("a numeral inside a word of Han characters", got)
	}
	if got = scanOf("삼일").hanNumerals(1); got != nil {
		t.Fatal(got)
	}
	if got = scanOf("十").hanNumerals(0); got != nil {
		t.Fatal(got)
	}
}
