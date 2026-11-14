//ff:func feature=combo type=helper control=iteration dimension=1
//ff:what Line — a combination of candidates printed on one line: its rank, its verdict and preference as nlgeul token writes them, and the candidate picked for each contest as nlgeul token writes a candidate, parted by " | "
package combo

import (
	"fmt"
	"strings"

	"github.com/geul-org/nlgeul/pkg/token"
)

// A candidate may hold a space (a boundary inside the contest), so the
// contests are parted by " | ".
func Line(rank int, h Combination, choices []token.Contest) string {
	picked := Picked(choices, h.Number)
	candidates := make([]string, 0, len(picked))
	for _, br := range picked {
		candidates = append(candidates, br.String())
	}
	return fmt.Sprintf("%d\t성립 %.2f 선호 %.2f\t%s", rank, h.Score.Verdict, h.Score.Pref, strings.Join(candidates, " | "))
}
