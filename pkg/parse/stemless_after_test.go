//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestStemlessAfter — a whole form right after a middle dot has no stem, a slot or another form; a slot right after a sign of a pair (closing or opening) or a punctuation mark has none; another form (a particle, the copula's) right after such a sign has one; a slot right after a Hangul syllable or a Latin letter has one; right after a digit a slot and a closing have none, but a closing that stands on the copula only, and a particle, have one
package parse

import "testing"

func TestStemlessAfter(t *testing.T) {
	slot := &entry{name: "선어말:였", form: "였", class: classSlot}
	particle := &entry{name: "목적격:을", form: "을", class: classParticle}
	closing := &entry{name: "맺음:면", form: "면", class: classClosing}
	copulaOnly := &entry{name: "맺음:야", form: "야", class: classClosing, needs: true}
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
		{slot, '3', true},
		{closing, '3', true},
		{closing, '하', false},
		{copulaOnly, '3', false},
		{particle, '3', false},
		{slot, 'A', false},
	} {
		if got := stemlessAfter(c.e, c.prev); got != c.want {
			t.Errorf("stemlessAfter(%s, %q) = %v", c.e.name, c.prev, got)
		}
	}
}
