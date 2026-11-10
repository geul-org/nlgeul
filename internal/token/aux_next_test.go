//ff:func feature=token type=reader control=sequence
//ff:what TestAuxNext — the letters from a place begin with an auxiliary verb's stem or do not
package token

import "testing"

func TestAuxNext(t *testing.T) {
	sc := scanOf("가버리다")
	if !sc.auxNext(1) || sc.auxNext(2) || sc.auxNext(4) {
		t.Fatal("버리 at 1 only")
	}
}
