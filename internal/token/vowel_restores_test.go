//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestVowelRestores — 해 comes from 하+여; 봐 줘 져 돼 from 보 주 지 되; 가 and 서 from themselves or from a dropped ㅡ; 내 and 세 from themselves
package token

import "testing"

func TestVowelRestores(t *testing.T) {
	cases := []struct {
		r    rune
		want string
	}{
		{'해', "하여"}, {'봐', "보아"}, {'줘', "주어"}, {'져', "지어"}, {'돼', "되어"},
		{'가', "가아그아"}, {'써', "써어쓰어"}, {'내', "내어"}, {'세', "세어"},
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
}
