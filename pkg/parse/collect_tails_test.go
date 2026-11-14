//ff:func feature=parse type=reader control=sequence
//ff:what TestCollectTails — 먹었다 has one tail; 분해한다 two (ㄴ다 and 다); before 요 the tail is a syllable longer and 요-less closings are skipped
package parse

import "testing"

func TestCollectTails(t *testing.T) {
	got := collectTails([]rune("먹었다"), 3, false)
	if len(got) != 1 || got[0].size != 4 || got[0].kind != kindFinal {
		t.Fatal(got)
	}
	if got = collectTails([]rune("분해한다"), 4, false); len(got) != 2 {
		t.Fatal(got)
	}
	if got = collectTails([]rune("먹어요"), 2, true); len(got) != 1 || got[0].size != 4 || got[0].marked != "먹+어+요" {
		t.Fatal(got)
	}
	if got = collectTails([]rune("먹던요"), 2, true); len(got) != 0 {
		t.Fatal(got)
	}
}
