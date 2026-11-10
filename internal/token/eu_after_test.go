//ff:func feature=token type=helper control=sequence
//ff:what TestEuAfter — a form with 으 in it that the table attaches after a consonant attaches after a consonant other than ㄹ; another form, and another condition, stay as the table has them
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEuAfter(t *testing.T) {
	if euAfter("을", afterBatchim) != rule.AfterHard || euAfter("은데", afterBatchim) != rule.AfterHard || euAfter("읍시다", afterBatchim) != rule.AfterHard {
		t.Fatal("forms with 으")
	}
	if euAfter("습니다", afterBatchim) != afterBatchim || euAfter("은", afterAny) != afterAny || euAfter("고", afterAny) != afterAny {
		t.Fatal("other forms and conditions")
	}
}
