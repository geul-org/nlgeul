//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadWordSets — this layer's own closed sets of words read from the lines of the closed-word table (set, word, and the mark only pkg/parse reads): the words of each set given in the order of their lines, every set given with a list (empty when no line names it); the lines of the sets pkg/parse reads (ClosedWords) are left to it, a set of neither is an error, and so is a mark on a line of a set given
package token

import (
	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/parse"
)

func loadWordSets(rows []data.Row, sets ...string) (map[string][]string, error) {
	names := map[string]int{}
	for set := range parse.ClosedWords() {
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
		if own == 0 {
			continue
		}
		if row.Opt(2) != "" {
			return nil, row.Errf(2, "이 집합의 줄에는 -를 적는다")
		}
		out[row.Cells[0]] = append(out[row.Cells[0]], row.Cells[1])
	}
	return out, nil
}
