//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestPrefinals — the slots before a closing: 잡히셨겠더 (fused ㅆ stops), 먹었었, 가시, none
package mini

import "testing"

func TestPrefinals(t *testing.T) {
	cases := []struct {
		word      string
		start     int
		stem, add int
	}{
		{"잡히셨겠더군", 5, 3, 5},
		{"먹었었다", 3, 1, 4},
		{"가시다", 2, 1, 2},
		{"예쁘다", 2, 2, 0},
		{"겠다", 1, 0, 2},
		{"다", 0, 0, 0},
	}
	for _, c := range cases {
		stem, size, fused := prefinals([]rune(c.word), c.start, 0)
		if stem != c.stem || size != c.add || fused != (c.word == "잡히셨겠더군") {
			t.Errorf("%s: stem %d size %d", c.word, stem, size)
		}
	}
}
