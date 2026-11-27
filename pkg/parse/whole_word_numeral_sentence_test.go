//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestWholeWordNumeralSentence — end to end over a small word table with 여러분 and 조사관: 여러분은 and 여러분들을 are read first as the pronoun (여러분[열쇠:체언]) and 조사관들은 as the noun, where the numeral 여러 or 조+사 and a unit came first before; the numeral and the unit stay a candidate
package parse

import (
	"strings"
	"testing"
)

func TestWholeWordNumeralSentence(t *testing.T) {
	r := streamReader(t, "words-whole-numeral.tsv")
	for _, c := range []struct{ in, want string }{
		{"여러분은", "여러분[열쇠:체언]+은[보조사:은]"},
		{"여러분들을", "여러분[열쇠:체언]+들[접미:들]+을[목적격:을]"},
		{"조사관들은", "조사관[열쇠:체언]+들[접미:들]+은[보조사:은]"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		if got := " │ " + strings.Join(words, " │ ") + " │ "; !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	if !netHas(r, "여러분은", "여러[수]+분[단위:분]+은[보조사:은]") {
		t.Fatal("the numeral and the unit are gone")
	}
}
