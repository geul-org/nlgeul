//ff:func feature=token type=rule control=sequence
//ff:what TestBoundOf — a bound stem with a word written wrong is read as the right form of the whole (입력/역활 → 입력/역할) when typos is true, judged as a misspelt form, below the same stem written right; not when typos is false (then nothing of it stands); a stem written right comes back as it is either way; bound is boundOf with no typos
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundOf(t *testing.T) {
	sc := typoScan("x", layerCost)
	right, candidates := sc.boundOf("입력/역활", rule.TailCase, true)
	_, plain := sc.boundOf("입력/역할", rule.TailCase, true)
	if right != "입력/역할" || len(candidates) != 1 || candidates[0].name != "대비:/" || candidates[0].v <= 0 || candidates[0].v >= plain[0].v {
		t.Fatal(right, candidates, plain)
	}
	if got, candidates := sc.boundOf("입력/역활", rule.TailCase, false); got != "입력/역활" || len(candidates) != 0 {
		t.Fatal(got, candidates)
	}
	if got, candidates := sc.boundOf("입력/출력", rule.TailCase, true); got != "입력/출력" || len(candidates) != 1 || candidates[0].name != "대비:/" || candidates[0].v <= 0 {
		t.Fatal(got, candidates)
	}
	if got, candidates := sc.boundOf("입력/역활", rule.TailPred, true); got != "입력/역활" || len(candidates) != 0 {
		t.Fatal("the last word must take what follows", got, candidates)
	}
	if got, candidates := sc.boundOf("입력/없는말", rule.TailCase, true); got != "입력/없는말" || len(candidates) != 0 {
		t.Fatal(got, candidates)
	}
	if candidates := sc.bound("입력/역활", rule.TailCase); len(candidates) != 0 {
		t.Fatal("bound reads no word written wrong", candidates)
	}
	if got, candidates := sc.boundOf("입력", rule.TailCase, true); got != "입력" || len(candidates) != 0 {
		t.Fatal("a stem with no bound sign is not a bound stem", got, candidates)
	}
}
