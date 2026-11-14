//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestLinkOf — each candidate gets its kind: a suffix, a vocative, 의, a case particle, 께서, 요, a shortened topic form, a directional, a comparing and a company particle, a range particle, an adverbial one, and a topic particle
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestLinkOf(t *testing.T) {
	cases := []struct {
		label, form string
		slot, want  int
	}{
		{"접미", "들", 0, rule.LinkSuffix}, {"호격", "야", 4, rule.LinkVoc}, {"관형격", "의", 3, rule.LinkAdn},
		{"주격", "가", 3, rule.LinkCase}, {"목적격", "를", 3, rule.LinkCase}, {"주격", "께서", 3, rule.LinkHonor},
		{"주격", "에서", 1, rule.LinkCase}, {"부사격", "에서", 1, rule.LinkAdv}, {"보조사", "요", 5, rule.LinkYo},
		{"부사격", "엔", 4, rule.LinkTopic}, {"부사격", "으로", 1, rule.LinkDir}, {"부사격", "처럼", 1, rule.LinkComp},
		{"부사격", "와", 1, rule.LinkWith}, {"보조사", "까지", 2, rule.LinkRange}, {"부사격", "대로", 2, rule.LinkRange},
		{"부사격", "토록", 2, rule.LinkAdv}, {"보조사", "는", 4, rule.LinkTopic},
	}
	for _, c := range cases {
		if got := linkOf(c.label, c.form, c.slot); got != c.want {
			t.Errorf("%s:%s = %d", c.label, c.form, got)
		}
	}
}
