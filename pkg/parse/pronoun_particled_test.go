//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestPronounParticled — the letters of a pronoun of the closed sets with a particle of the particle tables are pronoun-particled (나는, 이를, 우리가, 그를, 저는, 너희와); a pronoun alone, a pronoun with letters that are no particle, and other words are not (나, 이르, 나비, 이것, 검은, 우리다)
package parse

import "testing"

func TestPronounParticled(t *testing.T) {
	for _, c := range []struct {
		text string
		want bool
	}{
		{"나는", true},
		{"이를", true},
		{"우리가", true},
		{"그를", true},
		{"저는", true},
		{"너희와", true},
		{"나", false},
		{"이르", false},
		{"나비", false},
		{"이것", false},
		{"검은", false},
		{"우리다", false},
	} {
		if got := pronounParticled(c.text); got != c.want {
			t.Errorf("pronounParticled(%q) = %v, want %v", c.text, got, c.want)
		}
	}
}
