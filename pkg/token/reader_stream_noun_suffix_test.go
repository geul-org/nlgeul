//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamNounSuffix — end to end with the small table of testdata: 률 right after a nominal of the table that ends in a final consonant other than ㄴ is read as the suffix 접미:률 and comes first (보장률은, 등락률을, 분양률을, 불량률을, 보장률이다 with the copula after the suffix, 보장률); where the stem and 률 together are a word of the table (합격률, 확률, 법률, 능률), after ㄴ (출산률) or a vowel (증가률), and in words with no 률, it is not read and the first candidates stay; with no word table it is never read
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamNounSuffix(t *testing.T) {
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
		{"보장률은 높다", "보장[열쇠:체언]+률[접미:률]+은[보조사:은]", true},
		{"등락률을 봤다", "등락[열쇠:체언]+률[접미:률]+을[목적격:을]", true},
		{"분양률을 올렸다", "분양[열쇠:체언]+률[접미:률]+을[목적격:을]", true},
		{"불량률을 줄였다", "불량[열쇠:체언]+률[접미:률]+을[목적격:을]", true},
		{"보장률이다", "보장[열쇠:체언]+률[접미:률]+이[서술격:이]+다[맺음:다]", true},
		{"보장률", "보장[열쇠:체언]+률[접미:률]", true},
		{"합격률이 높다", "합격률[열쇠:체언]+이[주격:이]", false},
		{"확률이 높다", "확률[열쇠:체언]+이[주격:이]", false},
		{"법률을 지킨다", "법률[열쇠:체언]+을[목적격:을]", false},
		{"능률이 오른다", "능률[열쇠:체언]+이[주격:이]", false},
		{"출산률이 낮다", "출산[열쇠:체언] 률+이[주격:이]", false},
		{"증가률은 틀렸다", "증가[열쇠:체언] 률+은[보조사:은]", false},
		{"비율이 높다", "비율[열쇠:체언]+이[주격:이]", false},
		{"학교가 높다", "학교[열쇠:체언]+가[주격:가]", false},
		{"야투율은 낮다", "", false},
		{"피서객들이 몰렸다", "", false},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(with, c.sentence), "\n") + "\n"
		if c.first != "" && !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(with, c.sentence); strings.Contains(all, "[접미:률]") != c.suffix {
			t.Errorf("%s: 접미:률 read %v, want %v\n%s", c.sentence, !c.suffix, c.suffix, all)
		}
		if bare := streamText(without, c.sentence); strings.Contains(bare, "접미:률") {
			t.Errorf("%s: 접미:률 read with no word table:\n%s", c.sentence, bare)
		}
	}
}
