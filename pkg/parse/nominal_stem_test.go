//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNominalStem — a nominal stem: the 체언 of the word table, a dependent noun, a pronoun, a number (a decimal number too) or a new word with no name; not a verb, a closing or a part that is no stem
package parse

import "testing"

func TestNominalStem(t *testing.T) {
	yes := []Part{
		{Stem: true, Text: "뚜루뚜"},
		{Stem: true, Text: "씨", Name: Name{"열쇠", "체언"}},
		{Stem: true, Text: "거", Name: Name{"의존명사", "거"}},
		{Stem: true, Text: "나", Name: Name{Kind: "대명사"}},
		{Stem: true, Text: "3", Name: Name{Kind: "수"}},
		{Stem: true, Text: "3.5", Name: Name{Kind: "소수점"}},
		{Stem: true, Text: "①", Name: Name{"원", "수"}},
		{Stem: true, Text: "⑴", Name: Name{"괄호", "수"}},
		{Stem: true, Text: "n", Name: Name{"특수", "로마자"}},
	}
	no := []Part{
		{Text: "씨", Name: Name{"열쇠", "체언"}},
		{Stem: true, Text: "가", Name: Name{"열쇠", "동사"}},
		{Stem: true, Text: "다", Name: Name{"맺음", "다"}},
		{Stem: true, Text: "㈎", Name: Name{"괄호", "한글"}},
	}
	for _, p := range yes {
		if !nominalStem(p) {
			t.Fatal(p)
		}
	}
	for _, p := range no {
		if nominalStem(p) {
			t.Fatal(p)
		}
	}
}
