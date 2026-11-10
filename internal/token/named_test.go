//ff:func feature=token type=reader control=sequence
//ff:what TestNamed — an entry is found under its key by its name; a name that is not there gives an entry with no count
package token

import "testing"

func TestNamed(t *testing.T) {
	ix := scanOf("x").ix
	if e := ix.named('지', jiName); e.g == nil || e.form != "지" {
		t.Fatal(e)
	}
	if e := ix.named('지', "없음"); e.g != nil || e.count != 0 || e.name != "없음" {
		t.Fatal(e)
	}
}
