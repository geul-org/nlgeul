//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestRestores — a final consonant comes off (마신 → 시 + ㄴ, 합 → 하 + ㅂ); a fused ㅆ as 았/었/였 (갔 → 가 + 았 or 그 + 았, 했 → 하 + 였, 됐 → 되 + 었); a fused vowel as 아/어/여
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestRestores(t *testing.T) {
	cases := []struct {
		r    rune
		lead int
		want string
	}{
		{'신', rule.JongN, "시ㄴ"}, {'합', rule.JongB, "하ㅂ"}, {'갔', rule.JongSS, "가았그았"}, {'했', rule.JongSS, "하였"}, {'됐', rule.JongSS, "되었"}, {'봐', rule.LeadVowel, "보아"},
	}
	for _, c := range cases {
		got := ""
		for _, o := range restores(c.r, c.lead) {
			got += string(o.stem) + o.piece
		}
		if got != c.want {
			t.Errorf("%c: %s", c.r, got)
		}
	}
}
