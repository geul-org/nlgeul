//ff:func feature=token type=rule control=sequence
//ff:what TestAddHeads — a particle is kept under its first letter; an entry of another class is not
package token

import "testing"

func TestAddHeads(t *testing.T) {
	ix := &index{heads: map[rune][]string{}}
	ix.addHeads([]*entry{{form: "에서", class: classParticle}, {form: "다", class: classClosing}})
	if len(ix.heads) != 1 || ix.heads['에'][0] != "에서" {
		t.Fatal(ix.heads)
	}
}
