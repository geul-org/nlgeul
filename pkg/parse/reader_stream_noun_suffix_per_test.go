//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixPer — end to end with the small table of testdata: 당 right after a nominal of the table, at its word's end, is read as the suffix 접미:당 and comes first (팀당, 사람당); where the stem and 당 together are a word of the table (가구당, 정당), where a particle or the copula is written onto it (경로당에, 팀당이다), and in words with no 당, it is not read; with no word table it is never read after a nominal
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixPer(t *testing.T) {
	with, err := New(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	without, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence, first string
		suffix          bool
	}{
		{"팀당 비율이 높다", "팀[열쇠:체언]+당[접미:당]", true},
		{"사람당 보장률", "사람[열쇠:체언]+당[접미:당]", true},
		{"가구당 비율", "가구당[열쇠:체언]", false},
		{"정당이 높다", "정당[열쇠:체언]+이[주격:이]", false},
		{"경로당에 높다", "", false},
		{"팀당이다", "", false},
		{"학교가 높다", "학교[열쇠:체언]+가[주격:가]", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:당]") != c.suffix {
			t.Errorf("%s: 접미:당 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "접미:당") {
			t.Errorf("%s: 접미:당 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
