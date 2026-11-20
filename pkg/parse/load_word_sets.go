//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loadWordSets — this layer's own closed sets of words read from the lines of the closed-word table (set, word): the words of each set given in the order of their lines, every set given with a list (empty when no line names it); the lines of the sets ClosedWords gives are left to it, and a set of neither is an error
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
)

func loadWordSets(rows []data.Row, sets ...string) (map[string][]string, error) {
	names := map[string]int{}
	for set := range ClosedWords() {
		names[set] = 0
	}
	out := map[string][]string{}
	for _, set := range sets {
		names[set], out[set] = 1, nil
	}
	for _, row := range rows {
		own, err := row.Pick(0, names)
		if err != nil {
			return nil, err
		}
		if own == 1 {
			out[row.Cells[0]] = append(out[row.Cells[0]], row.Cells[1])
		}
	}
	return out, nil
}
