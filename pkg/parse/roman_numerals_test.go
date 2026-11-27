//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestRomanNumerals — a Roman numeral written in Latin capitals that is a whole run and ends at a place is a number before a unit, named 수:로마십진 (XIV before 세); with 제 right before it at a word's head it begins at 제 and is named 차례 first; one letter, X repeated, a part of a longer run and a run of other letters are not read
package parse

import "testing"

func TestRomanNumerals(t *testing.T) {
	got := scanOf("XIV세").romanNumerals(3)
	if len(got) != 1 || got[0].text != "XIV" || got[0].name != romanNumName || got[0].pref >= 0 {
		t.Fatal(got)
	}
	if got = scanOf("제II장").romanNumerals(3); len(got) != 2 || got[0].text != "제II" || got[0].name != "차례" || got[1].text != "제II" || got[1].name != romanNumName {
		t.Fatal(got)
	}
	if got = scanOf("곰제II장").romanNumerals(4); len(got) != 1 || got[0].text != "II" {
		t.Fatal("제 after a letter is no prefix", got)
	}
	for text, end := range map[string]int{"V세": 1, "XX동": 2, "AII호": 3, "CD장": 2, "ii호": 2, "3호": 1} {
		if got = scanOf(text).romanNumerals(end); got != nil {
			t.Errorf("%q: %v", text, got)
		}
	}
}
