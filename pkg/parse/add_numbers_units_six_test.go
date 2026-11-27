//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAddNumbersUnitsSix — end to end with a word table of testdata that knows the same letters as words (승, 패, 타, 푼, 품 and 정): the six units 타, 승, 패, 푼, 홉, 품 at the end of the unit table are read first right after a number (3승 2패, 5타, 2홉, 5푼, and 3품 after the noun 정 in 정3품), and the same letters with no number before them stay the words of the table (승 패)
package parse

import (
	"slices"
	"testing"
)

func TestAddNumbersUnitsSix(t *testing.T) {
	r := streamReader(t, "words-units-six.tsv")
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"3승 2패", []string{"{0 2} 3[수:십진]+승[단위:승]", "{2 4} 2[수:십진]+패[단위:패]"}},
		{"5타", []string{"{0 2} 5[수:십진]+타[단위:타]"}},
		{"2홉", []string{"{0 2} 2[수:십진]+홉[단위:홉]"}},
		{"5푼", []string{"{0 2} 5[수:십진]+푼[단위:푼]"}},
		{"정3품", []string{"{0 1} 정[열쇠:체언]", "{1 3} 3[수:십진]+품[단위:품]"}},
		{"승 패", []string{"{0 1} 승[열쇠:체언]", "{1 2} 패[열쇠:체언]"}},
	} {
		if got := flatNodes(r, c.in, true); !slices.Equal(got, c.want) {
			t.Errorf("%s: %q, want %q", c.in, got, c.want)
		}
	}
}
