//ff:func feature=token type=helper control=sequence
//ff:what TestAnyWrote — true when some piece of the tail is written wrong, false for none, for an empty list and for a piece written right
package token

import "testing"

func TestAnyWrote(t *testing.T) {
	right := piece{text: "게", names: "맺음:게"}
	wrong := piece{text: "(ㄹ)게", names: "맺음:ㄹ게", wrote: "ㄹ께"}
	if anyWrote(nil) || anyWrote([]piece{right, right}) {
		t.Fatal("no piece is written wrong")
	}
	if !anyWrote([]piece{right, wrong}) || !anyWrote([]piece{wrong}) {
		t.Fatal("a piece is written wrong")
	}
}
