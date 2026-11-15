//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingSurnameNot — article 48 with every rule on: the readings a family name of the list at a unit's head does not change stay first (안양에, 한양으로, 조양이, 고양이가, 장군이, 국군의, 이사장이: the family name before 양 or 군 may be a candidate, but the word stands first; 오늘); where the title is the head of another word (김양식장에: 양 and then 식, 한양부로: 양 and then 부) no candidate of the stream or of the old reading is named 성 at all
package token

import (
	"slices"
	"testing"
)

func TestSpacingSurnameNot(t *testing.T) {
	r := spacingReader(t, "words-surname.tsv", SpacingArticles...)
	for _, c := range []struct {
		text  string
		first []string
		none  bool
	}{
		{"안양에 갔다", []string{"{0 3} 안양[열쇠:체언]+에[부사격:에]"}, false},
		{"한양으로 갔다", []string{"{0 4} 한양[열쇠:체언]+으로[부사격:으로]"}, false},
		{"조양이 떴다", []string{"{0 3} 조양[열쇠:체언]+이[주격:이]"}, false},
		{"고양이가 운다", []string{"{0 4} 고양이[열쇠:체언]+가[주격:가]"}, false},
		{"장군이 왔다", []string{"{0 3} 장군[열쇠:체언]+이[주격:이]"}, false},
		{"국군의 날", []string{"{0 3} 국군[열쇠:체언]+의[관형격:의]"}, false},
		{"이사장이 왔다", []string{"{0 4} 이사장[열쇠:체언]+이[주격:이]"}, false},
		{"오늘 왔다", []string{"{0 2} 오늘[열쇠:체언]"}, true},
		{"김양식장에 갔다", []string{"{0 1} 김[열쇠:체언]", "{1 5} 양식장[열쇠:체언]+에[부사격:에]"}, true},
		{"한양부로 갔다", []string{"{0 2} 한양[열쇠:체언]", "{2 4} 부[열쇠:체언]+로[부사격:로]"}, true},
	} {
		got := flatNodes(r, c.text, true)
		if len(got) < len(c.first) || !slices.Equal(got[:len(c.first)], c.first) {
			t.Error(c.text, got)
		}
		if c.none && (netHas(r, c.text, "[성:") || unitHas(r, c.text, "[성:")) {
			t.Error("a family name is read", c.text)
		}
	}
}
