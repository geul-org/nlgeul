//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what Line — a hypothesis printed on one line: its rank, its verdict and preference as nlgeul token writes them, and the reading picked for each bundle as nlgeul token writes a reading, parted by " | "
package hypo

import (
	"fmt"
	"strings"

	"github.com/geul-org/nlgeul/internal/token"
)

// A reading may hold a space (a boundary inside the bundle), so the
// bundles are parted by " | ".
func Line(rank int, h Hypothesis, choices []token.Bundle) string {
	picked := Picked(choices, h.Number)
	readings := make([]string, 0, len(picked))
	for _, br := range picked {
		readings = append(readings, br.String())
	}
	return fmt.Sprintf("%d\t성립 %.2f 선호 %.2f\t%s", rank, h.Score.Verdict, h.Score.Pref, strings.Join(readings, " | "))
}
