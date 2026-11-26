//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixFirm — end to end with the small table of testdata: 사 right after a nominal of the table is read as the suffix 접미:사 and comes first, with particles and the copula after it (보장사는, 보장사로부터, 보장사이다); where the stem and 사 together are a word of the table (신문사, 의사), where they are two words of the table written together (기동전사: 기동 and 전사), after a word the table does not hold, and in words with no 사, it is not read; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixFirm(t *testing.T) {
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
		{"보장사는 높다", "보장[열쇠:체언]+사[접미:사]+는[보조사:는]", true},
		{"보장사로부터 높다", "보장[열쇠:체언]+사[접미:사]+로부터[부사격:로부터]", true},
		{"보장사이다", "보장[열쇠:체언]+사[접미:사]+이[서술격:이]+다[맺음:다]", true},
		{"신문사가 높다", "신문사[열쇠:체언]+가[주격:가]", false},
		{"의사가 높다", "의사[열쇠:체언]+가[주격:가]", false},
		{"기동전사가 높다", "", false},
		{"야투사는 낮다", "", false},
		{"학교가 높다", "학교[열쇠:체언]+가[주격:가]", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:사]") != c.suffix {
			t.Errorf("%s: 접미:사 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "접미:사") {
			t.Errorf("%s: 접미:사 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
