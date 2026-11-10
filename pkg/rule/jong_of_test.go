//ff:func feature=rule type=helper control=sequence
//ff:what TestJongOf — 가 has no final consonant, 간 has ㄴ, 갔 has ㅆ, a Latin letter is not Hangul
package rule

import "testing"

func TestJongOf(t *testing.T) {
	if JongOf('가') != 0 || JongOf('간') != JongN || JongOf('갔') != JongSS || JongOf('A') != -1 || JongOf(0xD7A4) != -1 {
		t.Fatal("JongOf")
	}
}
