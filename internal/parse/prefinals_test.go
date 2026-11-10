//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPrefinals — the slots before a closing: 잡히셨겠더 (fused ㅆ stops), 먹었었, 가시, none; a third 었 and a 겠 out of order are read as a fused ㅆ; 더 before 겠 is not read
package parse

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
		{"먹었었었다", 4, 2, 5},
		{"가겠었다", 3, 2, 3},
		{"가더겠다", 3, 2, 2},
		{"가시더다", 3, 1, 4},
	}
	for _, c := range cases {
		stem, size, fused := prefinals([]rune(c.word), c.start, 0)
		if stem != c.stem || size != c.add || fused != (size%2 == 1) {
			t.Errorf("%s: stem %d size %d", c.word, stem, size)
		}
	}
}
