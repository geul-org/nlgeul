//ff:func feature=parse type=helper control=sequence
//ff:what TestSinoNumerals — before 인 only the runs of Sino-Korean numerals are kept (삼, 십이), and a run of Han numerals (三), not a native numeral (네, 한, 몇), a Latin letter (N) or a Roman numeral (II); before any other unit all are kept
package parse

import "testing"

func TestSinoNumerals(t *testing.T) {
	words := []numeral{{"삼", -1, numName}, {"네", -1, numName}, {"십이", -2, numName}, {"한", -1, numName}, {"몇", -1, numName}, {"N", -1, numName}, {"三", -1, numName}, {"II", -1, romanNumName}}
	got := sinoNumerals("인", words)
	if len(got) != 3 || got[0].text != "삼" || got[1].text != "십이" || got[1].pref != -2 || got[2].text != "三" {
		t.Fatal(got)
	}
	if got := sinoNumerals("명", words); len(got) != len(words) {
		t.Fatal(got)
	}
	if got := sinoNumerals("인", nil); got != nil {
		t.Fatal(got)
	}
}
