//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunFinish — the end of a unit gives the stretches of the live states (the first touches the head of a unit that was not cut); when nothing reads what is left after a cut the letters are one stretch with no reading, and when nothing reads a unit that was not cut there is none
package token

import "testing"

func TestNetRunFinish(t *testing.T) {
	r, _ := windowRun(t, "words-far.tsv", windowText(1))
	r.window = false
	for r.j < r.n {
		r.step()
	}
	out := r.finish()
	if len(out) != 1 || out[0].contest.Text != windowText(1) || r.stats.unread != 0 {
		t.Fatal(out)
	}
	r.best[r.n], r.c = nil, 4
	out = r.finish()
	if len(out) != 1 || out[0].contest.Text != windowText(1)[len("고양이가"):] || r.stats.unread != 1 || out[0].cells.nodes != nil {
		t.Fatal(out)
	}
	r.c = 0
	if out = r.finish(); len(out) != 0 {
		t.Fatal("a unit nothing reads", out)
	}
}
