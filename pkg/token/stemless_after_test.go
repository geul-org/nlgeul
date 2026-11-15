//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestStemlessAfter — a whole form right after a middle dot has no stem, a slot or another form; a slot right after a sign of a pair (closing or opening) or a punctuation mark has none; another form (a particle, the copula's) right after such a sign has one; a slot right after a Hangul syllable, a digit or a Latin letter has one
package token

import "testing"

func TestStemlessAfter(t *testing.T) {
	slot := &entry{name: "선어말:였", form: "였", class: classSlot}
	particle := &entry{name: "목적격:을", form: "을", class: classParticle}
	for _, c := range []struct {
		e    *entry
		prev rune
		want bool
	}{
		{particle, '·', true},
		{slot, 'ㆍ', true},
		{slot, ')', true},
		{slot, '’', true},
		{slot, '\'', true},
		{slot, '(', true},
		{slot, ',', true},
		{slot, '/', true},
		{particle, ')', false},
		{particle, '/', false},
		{slot, '하', false},
		{slot, '3', false},
		{slot, 'A', false},
	} {
		if got := stemlessAfter(c.e, c.prev); got != c.want {
			t.Errorf("stemlessAfter(%s, %q) = %v", c.e.name, c.prev, got)
		}
	}
}
