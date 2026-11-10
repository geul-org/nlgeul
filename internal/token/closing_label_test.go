//ff:func feature=token type=reader control=sequence
//ff:what TestClosingLabel — 꾸밈 for a modifying closing, 이름 for a naming one, else 맺음
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestClosingLabel(t *testing.T) {
	if closingLabel(rule.KindAdn) != "꾸밈" || closingLabel(rule.KindNoun) != "이름" || closingLabel(rule.KindFinal|rule.KindLink) != "맺음" {
		t.Fatal("closingLabel")
	}
}
