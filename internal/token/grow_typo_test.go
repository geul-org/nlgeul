//ff:func feature=token type=reader control=sequence
//ff:what TestGrowTypo — a reading has letters written wrong when a word before or the word added has
package token

import "testing"

func TestGrowTypo(t *testing.T) {
	plain := grow(&path{verdict: 1}, step{sig: "곰", verdict: 1})
	if plain.typo {
		t.Fatal(plain)
	}
	if p := grow(plain, step{sig: "역할", verdict: 1, typo: true}); !p.typo {
		t.Fatal("the word added is written wrong")
	}
	wrong := grow(&path{verdict: 1}, step{sig: "역할", verdict: 1, typo: true})
	if p := grow(wrong, step{sig: "곰", verdict: 1}); !p.typo {
		t.Fatal("a word before is written wrong")
	}
}
