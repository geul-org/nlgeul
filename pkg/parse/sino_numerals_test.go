//ff:func feature=parse type=helper control=sequence
//ff:what TestSinoNumerals — before 인 only the runs of Sino-Korean numerals are kept (삼, 십이), not a native numeral (네, 한, 몇) or a Latin letter (N); before any other unit all are kept
package parse

import "testing"

func TestSinoNumerals(t *testing.T) {
	words := []numeral{{"삼", -1}, {"네", -1}, {"십이", -2}, {"한", -1}, {"몇", -1}, {"N", -1}}
	got := sinoNumerals("인", words)
	if len(got) != 2 || got[0].text != "삼" || got[1].text != "십이" || got[1].pref != -2 {
		t.Fatal(got)
	}
	if got := sinoNumerals("명", words); len(got) != len(words) {
		t.Fatal(got)
	}
	if got := sinoNumerals("인", nil); got != nil {
		t.Fatal(got)
	}
}
