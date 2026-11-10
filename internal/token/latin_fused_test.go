//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestLatinFused — a piece taken out of a syllable right after a Latin letter or digit inside the stem is dead (api과 as api고+아, v2과); alive are 하, 되 and 이 as the restored syllable (push해, merge된, api인데), a stem of one syllable after a boundary, a Hangul stem and a tail with nothing fused; end to end api과 has no api고+(아) and push해 keeps push하+(여)
package token

import (
	"strings"
	"testing"
)

func TestLatinFused(t *testing.T) {
	if !scanOf("api과").latinFused(0, tail{start: 4, base: '고'}) || !scanOf("v2과").latinFused(0, tail{start: 3, base: '고'}) {
		t.Fatal("a piece out of a syllable right after a Latin letter or a digit")
	}
	if scanOf("push해").latinFused(0, tail{start: 5, base: '하'}) || scanOf("merge된").latinFused(0, tail{start: 6, base: '되'}) || scanOf("api인데").latinFused(0, tail{start: 4, base: '이'}) {
		t.Fatal("하, 되 and 이 are let be")
	}
	if scanOf("api과").latinFused(3, tail{start: 4, base: '고'}) {
		t.Fatal("a stem of one syllable after a boundary")
	}
	if scanOf("사과").latinFused(0, tail{start: 2, base: '고'}) {
		t.Fatal("a Hangul stem")
	}
	if scanOf("api과").latinFused(0, tail{start: 4}) {
		t.Fatal("nothing fused")
	}
	for _, opt := range []Options{{Asleep: true}, {Words: "testdata/words.tsv", Asleep: true}} {
		r, _ := New(opt)
		if got := strings.Join(r.Lines("api과", false), "\n"); strings.Contains(got, "api고+(아)") || !strings.Contains(got, "api[열:영숫자]+과[부사격:과]") {
			t.Errorf("api과 (%q)\n%s", opt.Words, got)
		}
		if got := strings.Join(r.Lines("push해", false), "\n"); !strings.Contains(got, "push하+(여)[맺음:어(합침)]") {
			t.Errorf("push해 (%q)\n%s", opt.Words, got)
		}
	}
	if got := strings.Join(Lines("api과", false), "\n"); !strings.Contains(got, "api[열:영숫자] 고+(아)[맺음:어(합침)]") {
		t.Fatal("a stem of one syllable after the boundary where the script changes\n" + got)
	}
}
