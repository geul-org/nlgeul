//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestLeftMark — a quotation before 하다 is divided (가셨다고, 먹는다고); other text stays whole, even when it ends like a tail (지시, 먹고)
package parse

import "testing"

func TestLeftMark(t *testing.T) {
	cases := [][2]string{{"가셨다고", "가+시+었+다+고"}, {"먹는다고", "먹+는+다+고"}, {"지시", "지시"}, {"먹고", "먹고"}}
	for _, c := range cases {
		if got := leftMark([]rune(c[0])); got != c[1] {
			t.Errorf("%s: %s", c[0], got)
		}
	}
}
