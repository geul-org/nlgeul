//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestOnsetOfAny — a syllable begins with the initial consonant of the first syllable of one of the forms: 건, 걸, 겁, 게 and 것 as 거 does, 껄 as 꺼 does; not 너, not 껄 when 꺼 is not among the forms, not a letter that is no Hangul syllable, not with no forms
package parse

import "testing"

func TestOnsetOfAny(t *testing.T) {
	for _, c := range []struct {
		r     rune
		forms []string
		want  bool
	}{
		{'건', []string{"거", "것"}, true},
		{'걸', []string{"것"}, true},
		{'겁', []string{"거"}, true},
		{'게', []string{"거"}, true},
		{'껄', []string{"거", "꺼"}, true},
		{'너', []string{"거", "것"}, false},
		{'껄', []string{"거", "것"}, false},
		{'A', []string{"거"}, false},
		{'거', nil, false},
	} {
		if got := onsetOfAny(c.r, c.forms); got != c.want {
			t.Errorf("%c %v: %v, want %v", c.r, c.forms, got, c.want)
		}
	}
}
