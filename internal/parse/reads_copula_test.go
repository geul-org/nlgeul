//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReadsCopula — 이 before the tail and 였 not after 하 read the copula stem; 예요 holds it; 는 never reads it; a word that is only 이 does not
package parse

import "testing"

func TestReadsCopula(t *testing.T) {
	cases := []struct {
		word  string
		i     int
		flags int
		want  bool
		add   int
	}{
		{"방식이다", 3, 0, true, 2},
		{"높이는", 2, fNoIda, false, 0},
		{"학교예요", 2, fSelfIda, true, 0},
		{"문제였다", 2, 0, true, 0},
		{"하였다", 1, 0, false, 0},
		{"이다", 1, 0, false, 0},
		{"예쁘다", 2, 0, false, 0},
	}
	for _, c := range cases {
		got, add := readsCopula([]rune(c.word), c.i, ending{flags: c.flags})
		if got != c.want || add != c.add {
			t.Errorf("%s: %v %d", c.word, got, add)
		}
	}
}
