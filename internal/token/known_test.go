//ff:func feature=token type=rule control=sequence
//ff:what TestKnown — without a word table only closed-set words are known, whatever follows; in the first layer any word of the table; from the second only when its kind takes what follows
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestKnown(t *testing.T) {
	none := scanOf("x").ix
	if name, ok := none.known("것", rule.TailPred); !ok || name != "의존명사:것" {
		t.Fatal(name, ok)
	}
	if _, ok := none.known("고양이", rule.TailCase); ok {
		t.Fatal("known without a table")
	}
	one := wordsScan("x", layerKnown).ix
	if name, ok := one.known("마신", rule.TailPred); !ok || name != "열쇠:체언" {
		t.Fatal(name, ok)
	}
	two := wordsScan("x", layerKind).ix
	if _, ok := two.known("마신", rule.TailPred); ok {
		t.Fatal("a noun before a closing")
	}
	if name, ok := two.known("작", rule.TailPred); !ok || name != "열쇠:용언" {
		t.Fatal(name, ok)
	}
	if name, ok := two.known("작", rule.TailAux); !ok || name != "열쇠:체언" {
		t.Fatal(name, ok)
	}
	if name, ok := two.known("것", rule.TailCase); !ok || name != "의존명사:것" {
		t.Fatal(name, ok)
	}
	if _, ok := two.known("것", rule.TailPred); ok {
		t.Fatal("a dependent noun before a closing")
	}
	if _, ok := two.known("없는말", rule.TailBare); ok {
		t.Fatal("unknown")
	}
}
