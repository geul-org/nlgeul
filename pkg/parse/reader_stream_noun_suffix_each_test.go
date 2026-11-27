//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixEach — end to end with the small table of testdata: 별 right after a nominal of the table is read as the suffix 접미:별 and comes first, after a stem of two syllables or more and of one syllable too (사업자별로, 필지별, 시별로), with the noun 별 of the table kept among the candidates (별[열쇠:체언]); where the stem and 별 together are a word of the table (지역별) it is not read; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixEach(t *testing.T) {
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
		{"사업자별로 높다", "사업자[열쇠:체언]+별[접미:별]+로[부사격:로]", true},
		{"필지별 높다", "필지[열쇠:체언]+별[접미:별]", true},
		{"시별로 높다", "시[열쇠:체언]+별[접미:별]+로[부사격:로]", true},
		{"지역별이다", "지역별[열쇠:체언]+이[서술격:이]+다[맺음:다]", false},
		{"지역별로 높다", "지역별[열쇠:체언]+로[부사격:로]", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:별]") != c.suffix {
			t.Errorf("%s: 접미:별 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "[접미:별]") {
			t.Errorf("%s: 접미:별 read with no word table:\n%s", c.sentence, bare)
		}
	}
	if all := streamText(with, "사업자별로 높다"); !strings.Contains(all, "별[열쇠:체언]+로[부사격:로]") {
		t.Errorf("사업자별로: the reading with the noun 별 is gone\n%s", all)
	}
}
