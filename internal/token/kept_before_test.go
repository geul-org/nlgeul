//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestKeptBefore — what a word from b up to end grows: the candidates that end at b, or, when a cue shows the letters as a word that follows a word no table lists, those kept at b where a sign stands (the new word 뚜루뚜 as well); never for a tail that is a changed form (shown), a fused one (base) or a wrong spelling (as)
package token

import "testing"

func TestKeptBefore(t *testing.T) {
	sc := spacingScan("뚜루뚜겸총무", "45")
	best := make([][]*path, 7)
	best[0] = []*path{{verdict: 1}}
	plain := sc.keptBefore(best, 3, 4, tail{start: 4})
	if len(plain) == 0 || plain[0].last.word.Parts[0].Text != "뚜루뚜" {
		t.Fatal("a cue shows 겸", plain)
	}
	if got := sc.keptBefore(best, 3, 5, tail{start: 5}); len(got) != 0 {
		t.Fatal("no cue shows 겸총", got)
	}
	sc.floor = 4
	if got := sc.keptBefore(best, 3, 4, tail{start: 4}); got != nil {
		t.Fatal("before the first place", got)
	}
	sc.floor = 0
	for _, t2 := range []tail{{start: 4, shown: "돌"}, {start: 4, as: "x"}, {start: 4, base: '가'}} {
		if got := sc.keptBefore(best, 3, 4, t2); len(got) != 0 {
			t.Fatal("a changed form", t2, got)
		}
	}
}
