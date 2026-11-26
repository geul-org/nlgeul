//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamNumberSuffix — end to end with the small table of testdata: 사 right after a number is read as the suffix 접미:사 and comes first (3사는, 제2사를, 3사이다), after a number and its unit too (6개사의); with a word of the table after the number (3사단이) the first is that word; 률 is not read after a number (3률은); with no word table 사 is not read; the reading of the number written apart from 사 stays among the candidates
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamNumberSuffix(t *testing.T) {
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
		{"3사는 높다", "3[수]+사[접미:사]+는[보조사:는]", true},
		{"제2사를 보다", "제2[차례]+사[접미:사]+를[목적격:를]", true},
		{"3사이다", "3[수]+사[접미:사]+이[서술격:이]+다[맺음:다]", true},
		{"6개사의 학교", "6[수]+개[단위:개]+사[접미:사]+의[관형격:의]", true},
		{"3률은 높다", "", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:사]") != c.suffix || strings.Contains(all, "[접미:률]") {
			t.Errorf("%s: 접미:사 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "접미:") {
			t.Errorf("%s: a suffix read with no word table:\n%s", c.sentence, bare)
		}
	}
	if apart := streamText(with, "3사는 높다"); !strings.Contains(apart, "3[수]") || !strings.Contains(apart, "사[열쇠:체언]+는[보조사:는]") {
		t.Errorf("3사는: the reading apart from the number is gone\n%s", apart)
	}
	if first := strings.Join(bestLines(with, "3사단이 높다"), "\n"); !strings.Contains(first, "사단[열쇠:체언]+이[주격:이]") {
		t.Errorf("3사단이: the word of the table is not first\n%s", first)
	}
}
