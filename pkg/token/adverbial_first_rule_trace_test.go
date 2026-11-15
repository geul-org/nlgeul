//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAdverbialFirstRuleTrace — with a trace the rule writes a 통과 line (verdict 1.00) for each candidate it speaks of, the adverbial 에서 and 서 (학교에서 has the two), and none for the nominative, for the new word 학교에서, or for the particles with no 부사격 에서·서 in them; 학교에서는 has its two lines and the nominative that its chain refuses is not asked
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestAdverbialFirstRuleTrace(t *testing.T) {
	const name = "\t" + ruleAdverbialFirst + "\t1.00\t통과"
	cases := []struct {
		text string
		want []string
	}{
		{"학교에서", []string{"  · 학교[열쇠:체언]+에서[부사격:에서]" + name, "  · 학교에+서[부사격:서]" + name}},
		{"학교에서는", []string{"  · 학교[열쇠:체언]+에서[부사격:에서]+는[보조사:는]" + name, "  · 학교에+서[부사격:서]+는[보조사:는]" + name}},
		{"여기서", []string{"  · 여기[열쇠:체언]+서[부사격:서]" + name}},
	}
	r := streamReader(t, "words-adverbial.tsv")
	for _, c := range cases {
		st := r.StreamTrace(c.text)
		st.Next()
		got := adverbialFirstLines(st.Notes())
		if !slices.Equal(got, c.want) {
			t.Errorf("%s:\n%q\n%q", c.text, got, c.want)
		}
	}
	st := r.StreamTrace("학교에서")
	st.Next()
	for _, n := range adverbialFirstLines(st.Notes()) {
		if strings.Contains(n, "주격") || strings.Contains(n, "  · 학교에서\t") {
			t.Error(n)
		}
	}
	none := r.StreamTrace("학교")
	none.Next()
	if got := adverbialFirstLines(none.Notes()); got != nil {
		t.Errorf("a word with no marker is not asked: %q", got)
	}
}
