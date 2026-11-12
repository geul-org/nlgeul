//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNumberReads — a word whose stem is a run of the digits 0 and 1 only is read twice, the number first and the same word with its stem named 수:2진법 second (the markers after it and the verdict and preference unchanged); a number with another digit, a number joined by a sign, a number with a sign before it and a hexadecimal number are read once
package token

import "testing"

func TestNumberReads(t *testing.T) {
	read := func(text string, stemEnd int) []step {
		sc := scanOf(text)
		return sc.numberReads(sc.wordOf(0, len(sc.s), tail{start: stemEnd}), 0, stemEnd)
	}
	two := read("1010", 4)
	if len(two) != 2 || two[0].sig != "1010[수]" || two[1].sig != "1010[수:2진법]" || two[0].verdict != two[1].verdict || two[0].pref != two[1].pref || two[0].end != two[1].end {
		t.Fatalf("%+v", two)
	}
	if two[0].word.Parts[0].Name != (Name{"수", ""}) || two[1].word.Parts[0].Name != (Name{"수", "2진법"}) {
		t.Fatal("the first candidate is not changed by the second", two[0].word.Parts[0].Name)
	}
	for _, c := range []struct {
		text string
		end  int
	}{{"102", 3}, {"1,000", 5}, {"1.01", 4}, {"10:10", 5}, {"-11", 3}, {"0x11", 4}, {"abc", 3}} {
		if got := read(c.text, c.end); len(got) != 1 {
			t.Errorf("%s: %d candidates", c.text, len(got))
		}
	}
	sc := scanOf("10개")
	e := sc.wordOf(0, 3, tail{start: 2, pieces: []piece{{text: "개", names: "단위:개", verdict: 1, tail: 1}}})
	if got := sc.numberReads(e, 0, 2); len(got) != 2 || got[1].word.Parts[1].Name != (Name{"단위", "개"}) || got[1].sig != "10[수:2진법]+개[단위:개]" {
		t.Fatalf("%+v", got)
	}
}
