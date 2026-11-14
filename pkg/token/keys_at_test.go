//ff:func feature=token type=reader control=sequence
//ff:what TestKeysAt — a token is looked up as itself, and as a fused token under its final consonant (ㅁ also under ㄻ) or its vowel
package token

import "testing"

func TestKeysAt(t *testing.T) {
	if string(keysAt('간')) != "간ㄴ" || string(keysAt('가')) != "가ㅓ" || string(keysAt('각')) != "각" || string(keysAt('x')) != "x" || string(keysAt('삶')) != "삶ㅁ" {
		t.Fatal(string(keysAt('간')), string(keysAt('가')))
	}
}
