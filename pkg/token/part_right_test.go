//ff:func feature=token type=rule control=sequence
//ff:what TestPartRight — a word that stands as it is comes back as it is; a word written wrong comes back as the right form when typos is true, and as it is with 0 when it is not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPartRight(t *testing.T) {
	ix := typoScan("x", layerCost).ix
	if got, v := ix.partRight("입력", rule.TailCase, true); got != "입력" || v != ix.partVerdict("입력", rule.TailCase) || v <= 0 {
		t.Fatal(got, v)
	}
	if got, v := ix.partRight("역활", rule.TailCase, true); got != "역할" || v <= 0 {
		t.Fatal(got, v)
	}
	if got, v := ix.partRight("역활", rule.TailCase, false); got != "역활" || v != 0 {
		t.Fatal("a word written wrong is not read so unless typos", got, v)
	}
}
