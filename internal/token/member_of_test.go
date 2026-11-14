//ff:func feature=token type=helper control=sequence
//ff:what TestMemberOf — the test made from the listed ways of cutting says yes for a line of spans that is listed and no for one that is not
package token

import "testing"

func TestMemberOf(t *testing.T) {
	spans := []Span{{0, 2}, {2, 5}}
	in := memberOf(map[string]bool{lineKey(spans): true})
	if !in(spans) || in([]Span{{0, 3}, {3, 5}}) || in(nil) {
		t.Fatal("listed and not")
	}
}
