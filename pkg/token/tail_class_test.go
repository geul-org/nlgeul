//ff:func feature=token type=reader control=sequence
//ff:what TestTailClass — a closing only verbs take makes a verb's tail; else a predicate piece a predicate tail; else a case particle a case tail; else any particle an auxiliary tail; no piece is a bare stem; a closing verbs and 있다, 없다 and 계시다 take (느냐) makes a tail of its own, a predicate tail after the slot 었 or 겠 or 시, and not after 더
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTailClass(t *testing.T) {
	if tailClass([]piece{{tail: rule.TailPred}, {tail: rule.TailVerb}}) != rule.TailVerb || tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailPred}}) != rule.TailPred || tailClass([]piece{{tail: rule.TailAux}, {tail: rule.TailCase}}) != rule.TailCase || tailClass([]piece{{tail: rule.TailAux}}) != rule.TailAux || tailClass(nil) != rule.TailBare {
		t.Fatal("tailClass")
	}
	ask := piece{names: "맺음:느냐", tail: rule.TailExist}
	if tailClass([]piece{ask}) != rule.TailExist || tailClass([]piece{{names: "선어말:았", tail: rule.TailPred}, ask}) != rule.TailPred || tailClass([]piece{{names: "선어말:겠", tail: rule.TailPred}, ask}) != rule.TailPred || tailClass([]piece{{names: "이음:으", tail: rule.TailPred}, {names: "선어말:시", tail: rule.TailPred}, ask}) != rule.TailPred || tailClass([]piece{{names: "선어말:더", tail: rule.TailPred}, ask}) != rule.TailExist {
		t.Fatal("tailClass before 느냐")
	}
}
