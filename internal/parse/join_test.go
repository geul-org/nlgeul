//ff:func feature=parse type=reader control=sequence
//ff:what TestJoin — a shortened 것 and a joined noun are both tried; a word that holds neither is left as it was
package parse

import "testing"

func TestJoin(t *testing.T) {
	w := &word{}
	join(w, []rune("이게"))
	if w.chain.marked != "이+게" {
		t.Fatal(w.chain)
	}
	w = &word{}
	join(w, []rune("할수"))
	if w.chain.marked != "할+수" {
		t.Fatal(w.chain)
	}
	w = &word{chain: chainFact{size: 9}}
	join(w, []rune("이게"))
	join(w, []rune("사과"))
	if w.chain.size != 9 || w.tails != nil {
		t.Fatal(w.chain)
	}
}
