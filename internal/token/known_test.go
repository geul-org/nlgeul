//ff:func feature=token type=rule control=sequence
//ff:what TestKnown — without a word table only closed-set words are known, fully, when their set takes what follows (an auxiliary verb's stem not alone, a determiner not before a closing, a word of two sets as either), and 이, 그 and 저 before particles or the copula are named pronouns while alone, or with a table, they stay determiners; in the first layer any word of the table, fully; from the second by its kind's graph: not when no kind takes what follows, less when rarely seen or one syllable long; a closed-set word is named by its set only when it stands as that set's kind
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestKnown(t *testing.T) {
	none := scanOf("x").ix
	if name, v := knownIn(none, "것", rule.TailCase, false); v != 1 || name != "의존명사:것" {
		t.Fatal(name, v)
	}
	if _, v := knownIn(none, "것", rule.TailPred, false); v > 0 {
		t.Fatal("a dependent noun before a closing, without a table")
	}
	if _, v := knownIn(none, "보", rule.TailBare, false); v > 0 {
		t.Fatal("an auxiliary verb's stem alone, without a table")
	}
	if _, v := knownIn(none, "그", rule.TailPred, false); v > 0 {
		t.Fatal("a determiner before a closing, without a table")
	}
	if name, v := knownIn(none, "두", rule.TailPred, false); v != 1 || name != "관형사:두,보조 용언:두" {
		t.Fatal(name, v)
	}
	if name, v := knownIn(none, "그", rule.TailCase, false); v != 1 || name != "대명사:그" {
		t.Fatal("a determiner before a particle is let through, as the pronoun it is there", name, v)
	}
	if name, v := knownIn(none, "이", rule.TailAux, false); v != 1 || name != "대명사:이" {
		t.Fatal("이 before an auxiliary particle, without a table", name, v)
	}
	if name, v := knownIn(none, "저", rule.TailCopula, false); v != 1 || name != "대명사:저" {
		t.Fatal("저 before the copula, without a table", name, v)
	}
	if name, v := knownIn(none, "그", rule.TailBare, false); v != 1 || name != "관형사:그" {
		t.Fatal("그 alone stays a determiner, without a table", name, v)
	}
	if name, v := knownIn(none, "어떤", rule.TailCase, false); v != 1 || name != "관형사:어떤" {
		t.Fatal("a determiner that is no pronoun keeps its name before a particle", name, v)
	}
	if _, v := knownIn(none, "고양이", rule.TailCase, false); v != 0 {
		t.Fatal("known without a table")
	}
	one := wordsScan("x", layerKnown).ix
	if name, v := knownIn(one, "마신", rule.TailPred, false); v != 1 || name != "열쇠:체언" {
		t.Fatal(name, v)
	}
	two := wordsScan("x", layerKind).ix
	if _, v := knownIn(two, "마신", rule.TailPred, false); v != 0 {
		t.Fatal("a noun before a closing")
	}
	if name, v := knownIn(two, "작", rule.TailPred, false); v <= 0 || name != "열쇠:형용사" {
		t.Fatal(name, v)
	}
	if _, v := knownIn(two, "작", rule.TailVerb, false); v != 0 {
		t.Fatal("an adjective before a verb's closing")
	}
	if name, v := knownIn(two, "먹", rule.TailVerb, false); v <= 0 || name != "열쇠:동사" {
		t.Fatal(name, v)
	}
	if name, v := knownIn(two, "작", rule.TailAux, false); v <= 0 || v >= 1 || name != "열쇠:체언" {
		t.Fatal(name, v)
	}
	if name, v := knownIn(two, "것", rule.TailCase, false); v <= 0 || name != "의존명사:것" {
		t.Fatal(name, v)
	}
	if _, v := knownIn(two, "것", rule.TailPred, false); v != 0 {
		t.Fatal("a dependent noun before a closing")
	}
	if name, v := knownIn(two, "물", rule.TailBare, false); v <= 0 || name != "열쇠:체언" {
		t.Fatal(name, v)
	}
	if name, v := knownIn(two, "이", rule.TailCase, false); v <= 0 || name != "열쇠:체언" {
		t.Fatal("a determiner's letters before a particle stand as the table's noun and are named so", name, v)
	}
	if name, v := knownIn(two, "이", rule.TailBare, false); v <= 0 || name != "관형사:이" {
		t.Fatal("alone it is the determiner", name, v)
	}
	if name, v := knownIn(two, "그", rule.TailBare, false); v <= 0 || name != "관형사:그" {
		t.Fatal("with a table 그 alone is the determiner", name, v)
	}
	if name, v := knownIn(two, "그", rule.TailCase, false); v != 0 || name == "대명사:그" {
		t.Fatal("with a table 그 is not renamed a pronoun before a particle", name, v)
	}
	if name, v := knownIn(two, "두", rule.TailPred, false); v <= 0 || name != "보조 용언:두" {
		t.Fatal("a word of two sets is named by the set it stands as", name, v)
	}
	if _, v := knownIn(two, "없는말", rule.TailBare, false); v != 0 {
		t.Fatal("unknown")
	}
}
