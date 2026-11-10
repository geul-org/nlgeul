//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestVowelRestores — 해 comes from 하+여; 봐 줘 져 돼 from 보 주 지 되; 가 and 서 from themselves or from a dropped ㅡ; 내 and 세 from themselves; 놔 from 노+아 and 놓+아 (end to end 놔 is 놓+아 and 놨 is 놓+았); another syllable from itself
package token

import (
	"strings"
	"testing"
)

func TestVowelRestores(t *testing.T) {
	cases := []struct {
		r    rune
		want string
	}{
		{'해', "하여"}, {'봐', "보아"}, {'줘', "주어"}, {'져', "지어"}, {'돼', "되어"},
		{'가', "가아그아"}, {'써', "써어쓰어"}, {'내', "내어"}, {'세', "세어"},
		{'놔', "노아놓아"}, {'뵈', "뵈어"},
	}
	for _, c := range cases {
		got := ""
		for _, o := range vowelRestores(c.r) {
			got += string(o.stem) + o.piece
		}
		if got != c.want {
			t.Errorf("%c: %s", c.r, got)
		}
	}
	r, _ := New(Options{Words: "testdata/words.tsv", Asleep: true})
	if got := strings.Join(r.Lines("놔", false), "\n"); !strings.Contains(got, "놓[보조 용언:놓]+(아)[맺음:어(합침)]") {
		t.Fatal("놔\n" + got)
	}
	if got := strings.Join(r.Lines("놨다", false), "\n"); !strings.Contains(got, "놓[보조 용언:놓]+(았)[선어말:ㅆ]+다[맺음:다]") {
		t.Fatal("놨다\n" + got)
	}
}
