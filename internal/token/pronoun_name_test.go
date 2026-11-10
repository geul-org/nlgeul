//ff:func feature=token type=reader control=sequence
//ff:what TestPronounName — 그 before a particle or the copula is named a pronoun; alone it stays a determiner; a determiner that is no pronoun keeps its name, and so does another set's name of the same word
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestPronounName(t *testing.T) {
	if got := pronounName("관형사:그", "그", rule.TailCase); got != "대명사:그" {
		t.Fatal(got)
	}
	if got := pronounName("관형사:저", "저", rule.TailCopula); got != "대명사:저" {
		t.Fatal(got)
	}
	if got := pronounName("관형사:그", "그", rule.TailBare); got != "관형사:그" {
		t.Fatal("alone", got)
	}
	if got := pronounName("관형사:두|보조 용언:두", "두", rule.TailAux); got != "관형사:두|보조 용언:두" {
		t.Fatal("no pronoun", got)
	}
	if got := pronounName("의존명사:이|관형사:이", "이", rule.TailAux); got != "의존명사:이|대명사:이" {
		t.Fatal("another set's name", got)
	}
	if name, v := scanOf("그").ix.known("그", rule.TailCase); name != "대명사:그" || v <= 0 {
		t.Fatal(name, v)
	}
	if name, _ := scanOf("그").ix.known("그", rule.TailBare); name != "관형사:그" {
		t.Fatal(name)
	}
}
