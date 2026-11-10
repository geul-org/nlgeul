//ff:func feature=token type=reader control=sequence
//ff:what TestTailClass — a predicate piece makes a predicate tail; else a case particle a case tail; else any particle an auxiliary tail; no piece is a bare stem
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestTailClass(t *testing.T) {
	if tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailPred}}) != rule.TailPred || tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailCase}}) != rule.TailCase || tailClass([]piece{{tail: rule.TailAux}}) != rule.TailAux || tailClass(nil) != rule.TailBare {
		t.Fatal("tailClass")
	}
}
