//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunLambdaAt — the best lambda of a span is the first candidate's, as the stretch of the span would give it; it is kept while the same number of states stand at the span and worked out again when another number do
package token

import "testing"

func TestNetRunLambdaAt(t *testing.T) {
	r, _ := windowSteps(t, 5, 8)
	var sp Span
	var states []*path
	for _, s := range r.sc.registry().all {
		if s.start == 0 && s.last.end == 4 {
			sp, states = Span{0, 4}, append(states, s)
		}
	}
	if len(states) == 0 {
		t.Fatal("no state at 0-4")
	}
	want := lambdaOf(branchesOf(r.sc.ix, states[:1])[0])
	got := r.lambdaAt(sp, states[:1])
	if got != want || r.lam[sp].states != 1 {
		t.Fatal(got, want)
	}
	r.lam[sp] = spanLambda{states: 1, value: 42}
	if r.lambdaAt(sp, states[:1]) != 42 {
		t.Fatal("kept")
	}
	if len(states) > 1 && r.lambdaAt(sp, states) == 42 {
		t.Fatal("worked out again for another number of states")
	}
}
