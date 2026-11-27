//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixRight — end to end with the small table of testdata: 권 right after a nominal of two syllables or more is read as the suffix 접미:권 and comes first (환매권, 편성권을, 환매권이다), and after a number and its unit (10위권); the unit 권 that counts books stays (2권); where the stem and 권 together are a word of the table (소유권) or the stem is one syllable (법권) it is not read; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixRight(t *testing.T) {
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
		{"환매권 높다", "환매[열쇠:체언]+권[접미:권]", true},
		{"편성권을 높다", "편성[열쇠:체언]+권[접미:권]+을[목적격:을]", true},
		{"환매권이다", "환매[열쇠:체언]+권[접미:권]+이[서술격:이]+다[맺음:다]", true},
		{"10위권 높다", "10[수:십진]+위[단위:위]+권[접미:권]", true},
		{"2권 높다", "2[수:십진]+권[단위:권]", false},
		{"소유권이 높다", "소유권[열쇠:체언]+이[주격:이]", false},
		{"법권 높다", "", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:권]") != c.suffix {
			t.Errorf("%s: 접미:권 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "[접미:권]") {
			t.Errorf("%s: 접미:권 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
