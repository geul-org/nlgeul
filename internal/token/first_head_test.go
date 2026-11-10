//ff:func feature=token type=helper control=sequence
//ff:what TestFirstHead — a syllable that begins with ㄴ or ㅅ, 네, another syllable, and a letter that is no syllable
package token

import "testing"

func TestFirstHead(t *testing.T) {
	if firstHead('는') != headN || firstHead('시') != headS || firstHead('네') != headN|headNe || firstHead('고') != 0 || firstHead('x') != 0 {
		t.Fatal(firstHead('는'), firstHead('시'), firstHead('네'))
	}
}
