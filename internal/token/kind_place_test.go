//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestKindPlace — the place of each kind in kindOrder; 0 for a bit that is not one of them
package token

import "testing"

func TestKindPlace(t *testing.T) {
	for i, k := range kindOrder {
		if kindPlace(k) != i {
			t.Errorf("%d: %d", k, kindPlace(k))
		}
	}
	if kindPlace(1<<20) != 0 {
		t.Fatal("not a kind")
	}
}
