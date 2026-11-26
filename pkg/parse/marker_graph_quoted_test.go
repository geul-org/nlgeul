//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestMarkerGraphQuoted — the modifying closing 는 (quotedClosings) gets one more exception, which lifts "the letter before it is no Hangul syllable" right after a quotation mark that closes (간다"는), so it stands there and still does not after another letter that is no Hangul (push는); the modifying 은 of hangulBound has no such exception and stands after neither
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphQuoted(t *testing.T) {
	cases := []struct {
		e      entry
		edges  int
		quoted float64
	}{
		{entry{name: "꾸밈:는", form: "는", class: classClosing}, 4, 1},
		{entry{name: "꾸밈:은", form: "은", class: classClosing, after: rule.AfterBatchim}, 4, -1},
	}
	for _, c := range cases {
		e := c.e
		if got := markerGraph(&e, nil, nil, false); got != c.edges {
			t.Errorf("%s: %d edges, want %d", e.name, got, c.edges)
		}
		p := rule.Piece{Form: e.form, HasStem: true, Prev: '먹', Next: -1}
		if v := verdict(e.g, p); v != 1 {
			t.Errorf("%s after a Hangul syllable: %v", e.name, v)
		}
		p.Prev = 'h'
		if v := verdict(e.g, p); v != -1 {
			t.Errorf("%s after a Latin letter: %v", e.name, v)
		}
		p.Prev, p.Quoted = '"', true
		if v := verdict(e.g, p); v != c.quoted {
			t.Errorf("%s right after a quotation mark that closes: %v, want %v", e.name, v, c.quoted)
		}
	}
}
