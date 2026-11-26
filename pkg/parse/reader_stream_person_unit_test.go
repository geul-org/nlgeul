//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPersonUnit — end to end, with and without a word table: 인 right after a run of digits or a Sino-Korean numeral is read as the unit 단위:인 (3인이, 3인당 with the suffix after it; 삼인이 with a word table, as a lone Sino-Korean digit is no number without one), and the reading of the same letters as the copula's modifying form stays beside it (3인: 3+이+ㄴ); after a native numeral it is no unit (네인, 한인)
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestReaderStreamPersonUnit(t *testing.T) {
	with, err := New(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	without, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		reads    []string
		unit     bool
	}{
		{"3인이 높다", []string{"3[수]+인[단위:인]+이[주격:이]"}, true},
		{"3인당 비율", []string{"3[수]+인[단위:인]+당[접미:당]"}, true},
		{"비율이 3인 학교", []string{"3[수]+인[단위:인]", "3[수]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]"}, true},
		{"네인 학교", nil, false},
		{"한인 학교", nil, false},
	}
	if lone := streamText(with, "삼인이 높다"); !strings.Contains(lone, "\n삼[수]+인[단위:인]+이[주격:이]") {
		t.Errorf("삼인이: the unit after a Sino-Korean numeral is not read with a word table\n%s", lone)
	}
	for _, c := range cases {
		all := streamText(with, c.sentence)
		bare := streamText(without, c.sentence)
		missing := slices.ContainsFunc(c.reads, func(read string) bool {
			return !strings.Contains(all, "\n"+read) || !strings.Contains(bare, "\n"+read)
		})
		if missing || strings.Contains(all+bare, "[단위:인]") != c.unit {
			t.Errorf("%s: want %q read with and without a word table, 단위:인 %v\n%s\n%s", c.sentence, c.reads, c.unit, all, bare)
		}
	}
}
