//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamQuotedModifier — with no word table, the 는 written right onto a quotation mark that closes is read as the modifying closing of the quoting verb (+는[꾸밈:는]) beside the auxiliary particle (+는[보조사:는]), which stays, after a quotation mark of either kind, a corner bracket and a full stop inside the quotation; right after a closing round bracket only the particle is read, and 은 after a quotation mark is no modifying closing
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamQuotedModifier(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	const adn, aux = "+는[꾸밈:는]\t성립 1.00", "+는[보조사:는]\t성립 1.00"
	for _, c := range []struct {
		in   string
		want []string
		none string
	}{
		{"그는 \"곧 떠난다\"는 말을 남겼다.", []string{"\"[닫기:1]", aux, adn, "말을"}, ""},
		{"그는 ‘곧 떠났다’는 소문을 들었다.", []string{"’[닫기:1]", aux, adn, "소문을"}, ""},
		{"그는 「곧 떠나자」는 뜻을 밝혔다.", []string{"」[닫기:1]", aux, adn, "뜻을"}, ""},
		{"그는 “곧 떠난다.”는 말을 남겼다.", []string{".[마침표]", "”[닫기:1]", aux, adn, "말을"}, ""},
		{"그는 (곧 떠난다)는 말을 남겼다.", []string{")[닫기:1]", aux, "말을"}, "+는[꾸밈:는]"},
		{"그는 \"떠남\"은 말을 남겼다.", []string{"\"[닫기:1]", "+은[보조사:은]", "말을"}, "+은[꾸밈:은]"},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		if c.none != "" && missingLine(lines, []string{c.none}) == "" {
			t.Errorf("%q: %q is read, in\n%s", c.in, c.none, strings.Join(lines, "\n"))
		}
	}
}
