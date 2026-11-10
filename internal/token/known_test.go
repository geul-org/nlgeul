//ff:func feature=token type=rule control=sequence
//ff:what TestKnown — without a word table only closed-set words are known, fully, when their set takes what follows (an auxiliary verb's stem not alone, a determiner not before a closing, a word of two sets as either); in the first layer any word of the table, fully; from the second by its kind's graph: not when no kind takes what follows, less when rarely seen or one syllable long
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestKnown(t *testing.T) {
	none := scanOf("x").ix
	if name, v := none.known("것", rule.TailCase); v != 1 || name != "의존명사:것" {
		t.Fatal(name, v)
	}
	if _, v := none.known("것", rule.TailPred); v > 0 {
		t.Fatal("a dependent noun before a closing, without a table")
	}
	if _, v := none.known("보", rule.TailBare); v > 0 {
		t.Fatal("an auxiliary verb's stem alone, without a table")
	}
	if _, v := none.known("그", rule.TailPred); v > 0 {
		t.Fatal("a determiner before a closing, without a table")
	}
	if name, v := none.known("두", rule.TailPred); v != 1 || name != "관형사:두|보조 용언:두" {
		t.Fatal(name, v)
	}
	if _, v := none.known("그", rule.TailCase); v != 1 {
		t.Fatal("a determiner before a particle is let through")
	}
	if _, v := none.known("고양이", rule.TailCase); v != 0 {
		t.Fatal("known without a table")
	}
	one := wordsScan("x", layerKnown).ix
	if name, v := one.known("마신", rule.TailPred); v != 1 || name != "열쇠:체언" {
		t.Fatal(name, v)
	}
	two := wordsScan("x", layerKind).ix
	if _, v := two.known("마신", rule.TailPred); v != 0 {
		t.Fatal("a noun before a closing")
	}
	if name, v := two.known("작", rule.TailPred); v <= 0 || name != "열쇠:형용사" {
		t.Fatal(name, v)
	}
	if _, v := two.known("작", rule.TailVerb); v != 0 {
		t.Fatal("an adjective before a verb's closing")
	}
	if name, v := two.known("먹", rule.TailVerb); v <= 0 || name != "열쇠:동사" {
		t.Fatal(name, v)
	}
	if name, v := two.known("작", rule.TailAux); v <= 0 || v >= 1 || name != "열쇠:체언" {
		t.Fatal(name, v)
	}
	if name, v := two.known("것", rule.TailCase); v <= 0 || name != "의존명사:것" {
		t.Fatal(name, v)
	}
	if _, v := two.known("것", rule.TailPred); v != 0 {
		t.Fatal("a dependent noun before a closing")
	}
	if _, v := two.known("없는말", rule.TailBare); v != 0 {
		t.Fatal("unknown")
	}
}
