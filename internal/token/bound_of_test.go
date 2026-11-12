//ff:func feature=token type=rule control=sequence
//ff:what TestBoundOf — a bound stem with a word written wrong is read as the right form of the whole (입력/역활 → 입력/역할) when typos is true, judged as a misspelt form, below the same stem written right; not when typos is false (then nothing of it stands); a stem written right comes back as it is either way; bound is boundOf with no typos
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBoundOf(t *testing.T) {
	sc := typoScan("x", layerCost)
	right, name, v := sc.boundOf("입력/역활", rule.TailCase, true)
	_, _, plain := sc.boundOf("입력/역할", rule.TailCase, true)
	if right != "입력/역할" || name != "대비:/" || v <= 0 || v >= plain {
		t.Fatal(right, name, v, plain)
	}
	if got, name, v := sc.boundOf("입력/역활", rule.TailCase, false); got != "입력/역활" || name != "" || v != 0 {
		t.Fatal(got, name, v)
	}
	if got, name, v := sc.boundOf("입력/출력", rule.TailCase, true); got != "입력/출력" || name != "대비:/" || v <= 0 {
		t.Fatal(got, name, v)
	}
	if got, _, v := sc.boundOf("입력/역활", rule.TailPred, true); got != "입력/역활" || v != 0 {
		t.Fatal("the last word must take what follows", got, v)
	}
	if got, _, v := sc.boundOf("입력/없는말", rule.TailCase, true); got != "입력/없는말" || v != 0 {
		t.Fatal(got, v)
	}
	if name, v := sc.bound("입력/역활", rule.TailCase); name != "" || v != 0 {
		t.Fatal("bound reads no word written wrong", name, v)
	}
	if got, name, v := sc.boundOf("입력", rule.TailCase, true); got != "입력" || name != "" || v != 0 {
		t.Fatal("a stem with no bound sign is not a bound stem", got, name, v)
	}
}
