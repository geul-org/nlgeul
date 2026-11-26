//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixOpen — end to end with the small table of testdata: 율 right after a nominal of the table that ends in a vowel or ㄴ is read as the suffix 접미:율 and comes first (증가율은, 출산율이, 증가율이다 with the copula after the suffix, 출산율); where the stem and 율 together are a word of the table (비율), after another final consonant (보장율), after a word the table does not hold (야투율), and in words with no 율, it is not read; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixOpen(t *testing.T) {
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
		{"증가율은 높다", "증가[열쇠:체언]+율[접미:율]+은[보조사:은]", true},
		{"출산율이 낮다", "출산[열쇠:체언]+율[접미:율]+이[주격:이]", true},
		{"증가율이다", "증가[열쇠:체언]+율[접미:율]+이[서술격:이]+다[맺음:다]", true},
		{"출산율", "출산[열쇠:체언]+율[접미:율]", true},
		{"비율이 높다", "비율[열쇠:체언]+이[주격:이]", false},
		{"보장율은 틀렸다", "", false},
		{"야투율은 낮다", "", false},
		{"학교가 높다", "학교[열쇠:체언]+가[주격:가]", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:율]") != c.suffix {
			t.Errorf("%s: 접미:율 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "접미:율") {
			t.Errorf("%s: 접미:율 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
