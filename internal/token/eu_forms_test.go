//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEuForms — a form with 으 does not attach to a stem that ends in ㄹ (살을, 알은, 살으면 have no such reading of the regular stems), and it does attach where the ㄹ is what a ㄷ stem changed into (들은 = 듣+은, 들으면 = 듣+으면, 들으시고 = 듣+으+시+고)
package token

import (
	"strings"
	"testing"
)

func TestEuForms(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	for _, bad := range []string{"살을", "살은", "살으면", "알음", "들은", "들으시고"} {
		got := strings.Join(r.Lines(bad, false), "\n")
		if strings.Contains(got, "살[열쇠:동사]+을[꾸밈") || strings.Contains(got, "살[열쇠:동사]+은[꾸밈") || strings.Contains(got, "살[열쇠:동사]+으면") || strings.Contains(got, "+음[이름:음]") || strings.Contains(got, "들[") {
			t.Errorf("%s: a form with 으 after a stem that ends in ㄹ\n%s", bad, got)
		}
	}
	for in, want := range map[string]string{"들은": "듣[열쇠:동사]+은[꾸밈:은]", "들으면": "듣[열쇠:동사]+으면[맺음:으면]", "들으시고": "듣[열쇠:동사]+으[이음:으]+시[선어말:시]+고[맺음:고]", "살면": "살[열쇠:동사]+면[맺음:면]"} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
}
