//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAuxNext — the letters from a place begin with an auxiliary verb's stem, as it stands or with a piece fused into its last syllable (줘, 줬, 봐, 볼, 둬, 놔, 버려), or do not (조 is not 주, 버티 is not 버리); end to end a main verb's 어/아 that leaves no trace is read before 줘 and 볼, not before 조
package token

import (
	"strings"
	"testing"
)

func TestAuxNext(t *testing.T) {
	sc := scanOf("가버리다")
	if !sc.auxNext(1) || sc.auxNext(2) || sc.auxNext(4) {
		t.Fatal("버리 at 1 only")
	}
	for _, in := range []string{"빼줘", "빼줬다", "가봐", "가볼까", "빼둬", "빼놔", "가버려"} {
		if !scanOf(in).auxNext(1) {
			t.Errorf("%s: an auxiliary verb's stem with a piece fused into it", in)
		}
	}
	if scanOf("가조").auxNext(1) || scanOf("가버티다").auxNext(1) {
		t.Fatal("조 for 주, or a stem that shares only the first letters")
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	r.ix.lex["빼"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	if got := strings.Join(r.Lines("빼줘", false), "\n"); !strings.Contains(got, "빼[열쇠:동사]+(어)[맺음:어(합침)] 주[보조 용언:주]+(어)[맺음:어(합침)]") {
		t.Fatal("빼줘\n" + got)
	}
	if got := strings.Join(r.Lines("가볼까", false), "\n"); !strings.Contains(got, "가[열쇠:동사]+(아)[맺음:어(합침)] 보[보조 용언:보]+(ㄹ)까[맺음:ㄹ까]") {
		t.Fatal("가볼까\n" + got)
	}
	if got := strings.Join(r.Lines("가조", false), "\n"); strings.Contains(got, "가[열쇠:동사]+(아)") {
		t.Fatal("가조: an unseen 어/아 before a syllable that is no auxiliary verb's stem\n" + got)
	}
}
