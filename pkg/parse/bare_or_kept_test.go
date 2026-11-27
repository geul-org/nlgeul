//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestBareOrKept — a stem with no name and a stem named by a kept name (○○ 가림:원) are bare or kept; a part that is no stem and a stem named by the table of shapes (ⓐ 원:로마자) are not
package parse

import "testing"

func TestBareOrKept(t *testing.T) {
	for _, c := range []struct {
		p    Part
		want bool
	}{
		{Part{Text: "○○", Stem: true}, true},
		{Part{Text: "○○", Name: nameOf("가림:원"), Stem: true}, true},
		{Part{Text: "○○", Name: nameOf("가림:원")}, false},
		{Part{Text: "ⓐ", Name: nameOf("원:로마자"), Stem: true}, false},
	} {
		if got := bareOrKept(c.p); got != c.want {
			t.Errorf("bareOrKept(%+v) = %v", c.p, got)
		}
	}
}
