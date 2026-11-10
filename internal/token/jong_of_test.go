//ff:func feature=token type=reader control=sequence
//ff:what TestJongOf — 가 has no final consonant, 간 has ㄴ, a Latin letter is not Hangul
package token

import "testing"

func TestJongOf(t *testing.T) {
	if jongOf('가') != 0 || jongOf('간') != jongN || jongOf('A') != -1 || jongOf(0xD7A4) != -1 {
		t.Fatal("jongOf")
	}
}
