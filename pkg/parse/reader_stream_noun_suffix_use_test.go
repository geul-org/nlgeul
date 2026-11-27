//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffixUse — end to end with the small table of testdata: 용 right after a nominal of two syllables or more is read as the suffix 접미:용 and comes first, with particles and the copula after it (트랙터용, 사업용 부동산을, 파괴용으로, 공장용이다); right after a number (4용으로), after a number and its unit (4인용) and after a run of Latin letters and digits the table does not hold (OTS용) too; where the stem and 용 together are a word of the table (어린이용, 사용, 군용), where they are two words of it written together (초보자용: 초보 and 자용) and after a nominal of one syllable (법용) it is not read; with no word table it is never read
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffixUse(t *testing.T) {
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
		{"트랙터용 높다", "트랙터[열쇠:체언]+용[접미:용]", true},
		{"사업용 부동산을 보다", "사업[열쇠:체언]+용[접미:용]", true},
		{"파괴용으로 높다", "파괴[열쇠:체언]+용[접미:용]+으로[부사격:으로]", true},
		{"공장용이다", "공장[열쇠:체언]+용[접미:용]+이[서술격:이]+다[맺음:다]", true},
		{"4용으로 높다", "4[수]+용[접미:용]+으로[부사격:으로]", true},
		{"4인용 높다", "4[수]+인[단위:인]+용[접미:용]", true},
		{"OTS용 높다", "OTS[열:영숫자]+용[접미:용]", true},
		{"어린이용 높다", "어린이용[열쇠:체언]", false},
		{"사용이 높다", "사용[열쇠:체언]+이[주격:이]", false},
		{"군용이 높다", "군용[열쇠:체언]+이[주격:이]", false},
		{"초보자용 높다", "", false},
		{"법용 높다", "", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:용]") != c.suffix {
			t.Errorf("%s: 접미:용 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "[접미:용]") {
			t.Errorf("%s: 접미:용 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
