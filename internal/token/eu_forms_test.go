//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEuForms — a form with 으 other than the modifying 은 and 을 does not attach to a stem that ends in ㄹ (살으면, 알음 have no such candidate), the modifying 은 and 을 do attach to such a stem of the word table (살은, 살을) but stand lower than a closing of the tables after the same stem (살면), no 들 is read as a stem, and a form with 으 attaches where the ㄹ is what a ㄷ stem changed into (들은 = 듣+은, 들으면 = 듣+으면, 들으시고 = 듣+으+시+고)
package token

import (
	"strconv"
	"strings"
	"testing"
)

func TestEuForms(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	for _, bad := range []string{"살을", "살은", "살으면", "알음", "들은", "들으시고"} {
		got := strings.Join(r.Lines(bad, false), "\n")
		if strings.Contains(got, "살[열쇠:동사]+으면") || strings.Contains(got, "+음[이름:음]") || strings.Contains(got, "들[열쇠") {
			t.Errorf("%s: a form with 으 after a stem that ends in ㄹ\n%s", bad, got)
		}
	}
	verdict := func(in, reading string) float64 {
		for _, l := range r.Lines(in, false) {
			if rest, ok := strings.CutPrefix(l, "  "+reading+"\t성립 "); ok {
				v, _ := strconv.ParseFloat(strings.Fields(rest)[0], 64)
				return v
			}
		}
		return -1
	}
	regular := verdict("살면", "살[열쇠:동사]+면[맺음:면]")
	for in, want := range map[string]string{"살을": "살[열쇠:동사]+을[꾸밈:을]", "살은": "살[열쇠:동사]+은[꾸밈:은]"} {
		if v := verdict(in, want); v <= 0 || v >= regular {
			t.Errorf("%s: %s should stand below %s (%.2f), got %.2f\n%s", in, want, "살면", regular, v, strings.Join(r.Lines(in, false), "\n"))
		}
	}
	for in, want := range map[string]string{"들은": "듣[열쇠:동사]+은[꾸밈:은]", "들으면": "듣[열쇠:동사]+으면[맺음:으면]", "들으시고": "듣[열쇠:동사]+으[이음:으]+시[선어말:시]+고[맺음:고]", "살면": "살[열쇠:동사]+면[맺음:면]"} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
}
