//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestRestored — a dropped ㅡ put back where the vowel follows the syllable before it (따라, 치러, 써, 바빠) stands; it does not for a stem the table gives as 르 or 러 irregular (모라, 이러), for the wrong vowel (따러) and for 아 after a stem of one syllable (싸, also after a boundary); a piece taken out of a syllable right after a Latin letter or digit inside the stem does not stand (api과, v2과), but 하, 되 and 이 (push해, merge된, api인데), a stem of one syllable after a boundary, a Hangul stem and a tail with nothing fused do; end to end the live ones are read and the dead ones are not, api과 has no api고+(아) and push해 keeps push하+(여); the graph is told whether the syllable stands in the letters as it is: 하 right before the tail's start stays 하 (말하, no 말하+아) and is not put back, while 해 is, and a syllable of any other stem (가, 서) is put back either way; end to end 하지는, 말하지는 and 하서 have no 하+(아) reading while 해, 해서, 가서 , 했다 and 가 keep theirs, with and without a word table
package token

import (
	"strings"
	"testing"
)

func TestRestored(t *testing.T) {
	ws := wordsScan("따라", layerCost)
	if !ws.restored(0, tail{start: 2, base: '르'}) {
		t.Fatal("따라 = 따르+아: 아 after ㅏ")
	}
	if !ws.restored(0, tail{start: 2}) || !ws.restored(0, tail{start: 2, base: '라'}) {
		t.Fatal("a tail with no restored ㅡ")
	}
	if wordsScan("따러", layerCost).restored(0, tail{start: 2, base: '르'}) {
		t.Fatal("따러 as 따르+어: 어 after ㅏ")
	}
	if !wordsScan("치러", layerCost).restored(0, tail{start: 2, base: '르'}) || !wordsScan("바빠", layerCost).restored(0, tail{start: 2, base: '쁘'}) {
		t.Fatal("치르+어, 바쁘+아")
	}
	if wordsScan("모라", layerCost).restored(0, tail{start: 2, base: '르'}) || !scanOf("모라").restored(0, tail{start: 2, base: '르'}) {
		t.Fatal("모르 is 르 irregular in the table: no 모르+아; without a word table nothing tells it")
	}
	reo := wordsScan("이러", layerCost)
	reo.ix.lex["이르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}, irr: irrReo}
	if reo.restored(0, tail{start: 2, base: '르'}) {
		t.Fatal("a stem the table gives as 러 irregular")
	}
	if wordsScan("싸", layerCost).restored(0, tail{start: 1, base: '쓰'}) || !wordsScan("써", layerCost).restored(0, tail{start: 1, base: '쓰'}) {
		t.Fatal("a stem of one syllable takes 어 only")
	}
	if !wordsScan("곰써", layerCost).restored(1, tail{start: 2, base: '쓰'}) || wordsScan("곰싸", layerCost).restored(1, tail{start: 2, base: '쓰'}) {
		t.Fatal("a stem of one syllable after a boundary")
	}
	if !wordsScan("모로", layerCost).restored(0, tail{start: 2, base: '르'}) {
		t.Fatal("a syllable not written with ㅏ or ㅓ")
	}

	if scanOf("api과").restored(0, tail{start: 4, base: '고'}) || scanOf("v2과").restored(0, tail{start: 3, base: '고'}) {
		t.Fatal("a piece out of a syllable right after a Latin letter or a digit")
	}
	if !scanOf("push해").restored(0, tail{start: 5, base: '하'}) || !scanOf("merge된").restored(0, tail{start: 6, base: '되'}) || !scanOf("api인데").restored(0, tail{start: 4, base: '이'}) {
		t.Fatal("하, 되 and 이 are taken apart after a Latin letter too")
	}
	if !scanOf("api과").restored(3, tail{start: 4, base: '고'}) {
		t.Fatal("a stem of one syllable after a boundary")
	}
	if !scanOf("사과").restored(0, tail{start: 2, base: '고'}) {
		t.Fatal("a Hangul stem")
	}
	if !scanOf("api과").restored(0, tail{start: 4}) {
		t.Fatal("nothing fused")
	}

	if scanOf("말하").restored(0, tail{start: 2, base: '하'}) || scanOf("하").restored(0, tail{start: 1, base: '하'}) || scanOf("push하").restored(0, tail{start: 5, base: '하'}) {
		t.Fatal("하 in the letters hides no 어/아 that left no trace")
	}
	if !scanOf("말해").restored(0, tail{start: 2, base: '하'}) || !scanOf("해").restored(0, tail{start: 1, base: '하'}) {
		t.Fatal("해 holds 하+여: the piece left its trace")
	}
	if !scanOf("가").restored(0, tail{start: 1, base: '가'}) || !scanOf("일서").restored(0, tail{start: 2, base: '서'}) {
		t.Fatal("a syllable of another stem is put back whether or not it stands as it is")
	}

	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["따르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["치르"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["쓰"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["바쁘"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	for in, want := range map[string]string{"따라": "따르[열쇠:동사]+(아)", "치러": "치르[열쇠:동사]+(어)", "써": "쓰[열쇠:동사]+(어)", "바빠": "바쁘[열쇠:형용사]+(아)"} {
		if got := strings.Join(r.Lines(in, false), "\n"); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	for in, dead := range map[string]string{"모라": "모르[", "따러": "따르[", "싸": "쓰["} {
		if got := strings.Join(r.Lines(in, false), "\n"); strings.Contains(got, dead) {
			t.Errorf("%s: a dropped ㅡ put back where the stem does not drop it\n%s", in, got)
		}
	}
	for _, opt := range []Options{{}, {Words: "testdata/words.tsv"}} {
		r, _ := New(opt)
		if got := strings.Join(r.Lines("api과", false), "\n"); strings.Contains(got, "api고+(아)") || !strings.Contains(got, "api[열:영숫자]+과[부사격:과]") {
			t.Errorf("api과 (%q)\n%s", opt.Words, got)
		}
		if got := strings.Join(r.Lines("push해", false), "\n"); !strings.Contains(got, "push하+(여)[맺음:어(합침)]") {
			t.Errorf("push해 (%q)\n%s", opt.Words, got)
		}
	}
	plain, _ := New(Options{})
	lexed, _ := New(Options{Words: "testdata/words.tsv"})
	for _, c := range []struct {
		in, want string
		dead     bool
	}{
		{"하지는", "(아)", true}, {"말하지는", "(아)", true}, {"하서", "(아)", true},
		{"해", "하+(여)[맺음:어(합침)]", false}, {"해서", "하+(여)서[맺음:어서(합침)]", false}, {"했다", "하+(였)[선어말:ㅆ]+다[맺음:다]", false},
		{"가서", "가+(아)서[맺음:어서(합침)]", false}, {"가", "가+(아)[맺음:어(합침)]", false},
	} {
		got, withTable := strings.Join(plain.Lines(c.in, false), "\n"), strings.Join(lexed.Lines(c.in, false), "\n")
		wantTable := c.want
		if !c.dead {
			wantTable = strings.Replace(c.want, "+(", "[열쇠:동사]+(", 1)
		}
		if strings.Contains(got, c.want) == c.dead || strings.Contains(withTable, wantTable) == c.dead {
			t.Errorf("%s: 하 holds no 어/아 that leaves no trace, 해 and any other stem keep theirs (dead: %v)\n%s\n%s", c.in, c.dead, got, withTable)
		}
	}
	bare, _ := New(Options{})
	if got := strings.Join(bare.Lines("api과", false), "\n"); !strings.Contains(got, "api[열:영숫자] 고+(아)[맺음:어(합침)]") {
		t.Fatal("a stem of one syllable after the boundary where the script changes\n" + got)
	}
	if got := strings.Join(bare.Lines("따라", false), "\n"); !strings.Contains(got, "따르+(아)[맺음:어(합침)]") {
		t.Fatal("without a word table 따르+(아) stands\n" + got)
	}
	if got := strings.Join(bare.Lines("싸 써", false), "\n"); strings.Contains(got, "쓰+(아)") || !strings.Contains(got, "쓰+(어)[맺음:어(합침)]") {
		t.Fatal("without a word table too a stem of one syllable takes 어 only\n" + got)
	}
}
