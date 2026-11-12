//ff:func feature=token type=reader control=selection
//ff:what closingLabel — the candidate of a closing by its kind: 꾸밈, 이름, else 맺음
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func closingLabel(kind int) string {
	switch {
	case kind&rule.KindAdn != 0:
		return "꾸밈"
	case kind&rule.KindNoun != 0:
		return "이름"
	}
	return "맺음"
}
