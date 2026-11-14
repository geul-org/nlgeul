//ff:func feature=token type=helper control=sequence
//ff:what TestEuAfter — a form with 으 in it that the table attaches after a consonant attaches after a consonant other than ㄹ; another form, and another condition, stay as the table has them
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEuAfter(t *testing.T) {
	if euAfter("을", rule.AfterBatchim) != rule.AfterHard || euAfter("은데", rule.AfterBatchim) != rule.AfterHard || euAfter("읍시다", rule.AfterBatchim) != rule.AfterHard {
		t.Fatal("forms with 으")
	}
	if euAfter("습니다", rule.AfterBatchim) != rule.AfterBatchim || euAfter("은", rule.AfterAny) != rule.AfterAny || euAfter("고", rule.AfterAny) != rule.AfterAny {
		t.Fatal("other forms and conditions")
	}
}
