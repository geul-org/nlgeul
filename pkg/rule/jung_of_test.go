//ff:func feature=rule type=helper control=sequence
//ff:what TestJungOf — 가 is ㅏ(0), 봐 is ㅘ(9), 깁 is ㅣ(20), a digit is not Hangul
package rule

import "testing"

func TestJungOf(t *testing.T) {
	if JungOf('가') != 0 || JungOf('봐') != 9 || JungOf('깁') != 20 || JungOf('3') != -1 || JungOf(0xD7A4) != -1 {
		t.Fatal("JungOf")
	}
}
