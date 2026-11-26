//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestPoliteBefore — a slot in 옵 stands before a closing that begins with a consonant, one in 오 before one that begins with ㄴ, ㄹ, ㅁ or a vowel; neither before a letter that is no Hangul syllable
package parse

import "testing"

func TestPoliteBefore(t *testing.T) {
	for _, c := range []struct {
		last, r rune
		want    bool
	}{
		{'옵', '고', true}, {'옵', '니', true}, {'옵', '소', true}, {'옵', '아', false}, {'옵', 'a', false},
		{'오', '니', true}, {'오', '리', true}, {'오', '며', true}, {'오', '이', true}, {'오', '고', false}, {'오', '다', false}, {'오', '.', false},
	} {
		if got := politeBefore(c.last, c.r); got != c.want {
			t.Errorf("%c before %c: %v", c.last, c.r, got)
		}
	}
}
