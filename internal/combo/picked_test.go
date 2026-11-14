//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what TestPicked — number 0 picks the first candidate of every contest; the last contest's digit is the lowest; the last number picks the last candidate of every contest; the contests have different numbers of candidates
package combo

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/token"
)

func TestPicked(t *testing.T) {
	// Each candidate is told apart by its preference: contest*10 + candidate.
	choices := []token.Contest{
		{Text: "가", Branches: []token.Branch{{Pref: 10}, {Pref: 11}}},
		{Text: "나", Branches: []token.Branch{{Pref: 20}, {Pref: 21}, {Pref: 22}}},
		{Text: "다", Branches: []token.Branch{{Pref: 30}, {Pref: 31}}},
	}
	cases := []struct {
		number int
		want   [3]float64
	}{
		{0, [3]float64{10, 20, 30}},
		{1, [3]float64{10, 20, 31}},
		{2, [3]float64{10, 21, 30}},
		{3, [3]float64{10, 21, 31}},
		{5, [3]float64{10, 22, 31}},
		{6, [3]float64{11, 20, 30}},
		{7, [3]float64{11, 20, 31}},
		{10, [3]float64{11, 22, 30}},
		{11, [3]float64{11, 22, 31}},
	}
	for _, c := range cases {
		got := Picked(choices, c.number)
		if len(got) != 3 || [3]float64{got[0].Pref, got[1].Pref, got[2].Pref} != c.want {
			t.Errorf("number %d: %+v, want %v", c.number, got, c.want)
		}
	}
}
