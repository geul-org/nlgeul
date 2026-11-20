//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestGlossAt — a round bracket right at the place that holds Han characters alone and is closed tells how many (대곤(大棍): 2, 운용한다(行陣和睦): 4, a compatibility ideograph counts); no bracket there, a bracket left open, an empty one, one with a letter that is no Han character (Hangul, Latin, a comma, a space) and a full-width bracket tell 0
package parse

import "testing"

func TestGlossAt(t *testing.T) {
	for _, c := range []struct {
		text string
		at   int
		want int
	}{
		{"대곤(大棍)을", 2, 2},
		{"운용한다(行陣和睦).", 4, 4},
		{"(大)", 0, 1},
		{"악(樂)", 1, 1},
		{"대곤(大棍)을", 0, 0},
		{"대곤(大棍)을", 7, 0},
		{"대곤", 2, 0},
		{"대곤(大棍", 2, 0},
		{"대곤()", 2, 0},
		{"대곤(대곤)", 2, 0},
		{"대곤(大A)", 2, 0},
		{"대곤(大棍,)", 2, 0},
		{"대곤(大 棍)", 2, 0},
		{"대곤（大棍）", 2, 0},
	} {
		if got := glossAt([]rune(c.text), c.at); got != c.want {
			t.Errorf("glossAt(%q, %d) = %d, want %d", c.text, c.at, got, c.want)
		}
	}
}
