//ff:func feature=parse type=rule control=sequence
//ff:what TestJungOf — 가 is ㅏ(0), 봐 is ㅘ(9), 기 is ㅣ(20), a digit is not Hangul
package parse

import "testing"

func TestJungOf(t *testing.T) {
	if jungOf([]rune("가")[0]) != 0 || jungOf([]rune("봐")[0]) != 9 || jungOf([]rune("깁")[0]) != 20 || jungOf(0x33) != -1 || jungOf(0xD7A4) != -1 {
		t.Fatal("jungOf")
	}
}
