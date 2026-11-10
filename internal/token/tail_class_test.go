//ff:func feature=token type=reader control=sequence
//ff:what TestTailClass — a closing only verbs take makes a verb's tail; else a predicate piece a predicate tail; else a case particle a case tail; else any particle an auxiliary tail; no piece is a bare stem
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTailClass(t *testing.T) {
	if tailClass([]piece{{tail: rule.TailPred}, {tail: rule.TailVerb}}) != rule.TailVerb || tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailPred}}) != rule.TailPred || tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailCase}}) != rule.TailCase || tailClass([]piece{{tail: rule.TailAux}}) != rule.TailAux || tailClass(nil) != rule.TailBare {
		t.Fatal("tailClass")
	}
}
