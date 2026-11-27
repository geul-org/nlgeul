//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixSystem — end to end with the small table of testdata: 제 right after a nominal of two syllables or more is read as the suffix 접미:제 and comes first (연봉제를, 교대제, 수처리제로), and after a number and its unit (1차제); with a number right after 제 it is the prefix of that number (요추제1, 갑제13) and no suffix; the shortened 제 (제 생각, 제가) and the order 제1조 stay as they were; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixSystem(t *testing.T) {
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
		{"연봉제를 높다", "연봉[열쇠:체언]+제[접미:제]+를[목적격:를]", true},
		{"교대제 높다", "교대[열쇠:체언]+제[접미:제]", true},
		{"수처리제로 높다", "수처리[열쇠:체언]+제[접미:제]+로[부사격:로]", true},
		{"1차제 높다", "1[수]+차[단위:차]+제[접미:제]", true},
		{"요추제1 높다", "", false},
		{"갑제13 높다", "", false},
		{"제 생각 높다", "제[열쇠:체언]", false},
		{"제1조 높다", "제1[차례]+조[단위:조]", false},
		{"제가 높다", "제[열쇠:체언]+가[주격:가]", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first) {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:제]") != c.suffix {
			t.Errorf("%s: 접미:제 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "[접미:제]") {
			t.Errorf("%s: 접미:제 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
