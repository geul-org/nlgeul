//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestSaisiotOnset — 사이시옷 may come before a syllable that begins with ㄱ ㄷ ㅂ ㅅ ㅈ, ㄴ ㅁ, or ㅇ with ㅣ ㅑ ㅕ ㅛ ㅠ (잎, 일, 야, 열, 요, 윷), not before 이 itself, ㅇ with another vowel, ㅊ ㅋ ㅌ ㅍ, a doubled consonant, ㅎ or ㄹ, nor before what is not a Hangul syllable
package parse

import "testing"

func TestSaisiotOnset(t *testing.T) {
	for _, r := range "값대비살집내목잎일야열요윷" {
		if !saisiotOnset(r) {
			t.Errorf("%c: 사이시옷 may come before it", r)
		}
	}
	for _, r := range "이안웃얼애층카턱피깨떡빵쌀쪽함라A1" {
		if saisiotOnset(r) {
			t.Errorf("%c: no 사이시옷 before it", r)
		}
	}
}
