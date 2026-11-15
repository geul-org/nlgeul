//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestAloneWords — the words are the closed sets of bound nouns and determiners of the grammar tables, all of both (시 and 수 of the bound nouns, 이 and 그 of the determiners) and none of the other sets (the auxiliary verb 주, the conjunctive adverb 그러나)
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
)

func TestAloneWords(t *testing.T) {
	words := aloneWords()
	sets := parse.ClosedWords()
	if len(words) != len(sets["의존명사"])+len(sets["관형사"]) {
		t.Fatal(len(words))
	}
	for _, w := range []string{"시", "수", "이", "그"} {
		if !slices.Contains(words, w) {
			t.Errorf("%s is missing", w)
		}
	}
	for _, w := range []string{"주", "그러나"} {
		if slices.Contains(words, w) {
			t.Errorf("%s is of another set", w)
		}
	}
}
