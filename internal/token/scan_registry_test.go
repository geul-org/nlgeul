//ff:func feature=token type=helper control=sequence
//ff:what TestScanRegistry — a scan with none gets one that begins at 0, and the same one each time after
package token

import "testing"

func TestScanRegistry(t *testing.T) {
	sc := &scan{}
	r := sc.registry()
	if r == nil || r.next != 0 || sc.registry() != r {
		t.Fatal(r)
	}
	own := newRegistry(3)
	if (&scan{reg: own}).registry() != own {
		t.Fatal("the one it has")
	}
}
