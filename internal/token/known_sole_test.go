//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestKnownSole — with the word table of testdata every one-syllable stem of the table stands higher, under the same name, before a tail its kind takes when its word is the whole of its unit (하 and 가, seen enough, fully), and every longer stem stands the same; without a word table sole changes neither the name nor the verdict of a closed-set word, nor of a word no set has
package token

import (
	"testing"
	"unicode/utf8"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestKnownSole(t *testing.T) {
	fits := map[int]int{kindNominal: rule.TailCase, kindVerb: rule.TailPred, kindAdj: rule.TailAdj, kindAdverb: rule.TailAux, kindDeterminer: rule.TailBare, kindAlone: rule.TailBare}
	same := func(ix *index, stem string, tail int) bool {
		n, v := ix.known(stem, tail, false)
		m, w := ix.known(stem, tail, true)
		return n == m && v == w
	}
	ix := wordsScan("x", layerCost).ix
	ones := 0
	for stem, w := range ix.lex {
		k := fits[w.kinds&-w.kinds]
		name, in := ix.known(stem, k, false)
		sole, v := ix.known(stem, k, true)
		one := utf8.RuneCountInString(stem) == 1
		if in <= 0 || sole != name || (one && v <= in) || (!one && v != in) {
			t.Error(stem, name, in, sole, v)
		}
		if one {
			ones++
		}
	}
	if ones == 0 || ones == len(ix.lex) {
		t.Fatal("the table has stems of one syllable and of more", ones, len(ix.lex))
	}
	if _, v := ix.known("하", rule.TailVerb, true); v != 1 {
		t.Fatal("하 seen enough, as the whole of its unit", v)
	}
	if _, v := ix.known("가", rule.TailPred, true); v != 1 {
		t.Fatal("가 seen enough, as the whole of its unit", v)
	}
	none := scanOf("x").ix
	for stem := range none.words {
		if !same(none, stem, rule.TailCase) || !same(none, stem, rule.TailPred) || !same(none, stem, rule.TailBare) {
			t.Error("without a word table", stem)
		}
	}
	if _, v := none.known("고양이", rule.TailCase, true); none.layers != layerNone || len(none.words) == 0 || !same(none, "고양이", rule.TailCase) || v != 0 {
		t.Fatal("a word no set has, without a word table", v)
	}
}
