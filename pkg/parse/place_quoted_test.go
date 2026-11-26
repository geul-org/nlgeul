//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPlaceQuoted — a whole form right after a quotation mark that closes (quoteShuts: " ' ” ’ 」 』 ｣) is told so (Quoted), and one right after a closing bracket, an opening quotation mark or a Hangul syllable is not
package parse

import "testing"

func TestPlaceQuoted(t *testing.T) {
	adn := &entry{name: "꾸밈:는", form: "는", class: classClosing}
	for _, c := range []struct {
		in   string
		want bool
	}{
		{"간다\"는", true},
		{"간다'는", true},
		{"간다”는", true},
		{"간다’는", true},
		{"간다」는", true},
		{"간다』는", true},
		{"간다｣는", true},
		{"간다)는", false},
		{"간다“는", false},
		{"간다가는", false},
	} {
		n := len([]rune(c.in))
		start, p, ok := scanOf(c.in).place(adn, n, false)
		if !ok || start != n-1 || p.Quoted != c.want {
			t.Errorf("%q: Quoted %v, want %v (start %d, ok %v)", c.in, p.Quoted, c.want, start, ok)
		}
	}
	if _, p, ok := scanOf("는").place(adn, 1, false); !ok || p.Quoted {
		t.Error("a form at the head of its unit stands after no quotation mark", p, ok)
	}
}
