//ff:func feature=token type=rule control=sequence
//ff:what TestSetReads — without a word table a closed-set word is read once for each set it is in, with the verdict of its one graph; a pronoun's letters before a particle are named the pronoun; a word whose set does not take what follows, and a word of no set, have no candidate
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSetReads(t *testing.T) {
	ix := scanOf("x").ix
	got := ix.setReads(ix.words["두"], "두", rule.TailPred)
	if len(got) != 2 || got[0].name != "관형사:두" || got[1].name != "보조 용언:두" || got[0].v != 1 || got[1].v != 1 || got[0].kind != 0 {
		t.Fatal("a word of two sets is read twice, with one verdict", got)
	}
	if got = ix.setReads(ix.words["그"], "그", rule.TailCase); len(got) != 1 || got[0].name != "대명사:그" {
		t.Fatal("a pronoun before a particle", got)
	}
	if got = ix.setReads(ix.words["그"], "그", rule.TailPred); len(got) != 0 {
		t.Fatal("a determiner before a closing", got)
	}
	if got = ix.setReads(nil, "고양이", rule.TailBare); len(got) != 0 {
		t.Fatal("no set", got)
	}
}
