//ff:func feature=token type=rule control=sequence
//ff:what TestPartTypo — a word of a bound stem that the word table gives as a wrong way to write another word comes back as the right form with the verdict of the right form as a misspelt form (below a word written right); not when the right form does not take what follows, and a word the table does not give as misspelt comes back as it is with 0
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPartTypo(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	right, v := ix.partTypo("역활", rule.TailCase)
	if right != "역할" || v <= 0 || v >= ix.partVerdict("역할", rule.TailCase) {
		t.Fatal(right, v)
	}
	if got, v := ix.partTypo("맞기", rule.TailCase); got != "맞기" || v != 0 {
		t.Fatal("a misspelt verb stem stands with no particle", got, v)
	}
	if got, v := ix.partTypo("입력", rule.TailCase); got != "입력" || v != 0 {
		t.Fatal("a word that is not misspelt", got, v)
	}
}
