//ff:func feature=mini type=rule control=sequence
//ff:what TestJongOf — 가 has no final consonant, 간 has ㄴ, 갔 has ㅆ, a Latin letter is not Hangul
package mini

import "testing"

func TestJongOf(t *testing.T) {
	if jongOf(0xAC00) != 0 || jongOf([]rune("간")[0]) != jongN || jongOf([]rune("갔")[0]) != jongSS || jongOf(0x41) != -1 || jongOf(0xD7A4) != -1 {
		t.Fatal("jongOf")
	}
}
