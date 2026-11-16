//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loadClosed — the closed sets of words read from the lines of their table (set, word): the words of each of the three sets (auxiliary verbs, determiners, conjunctions) in the order of their lines; the lines of other sets, which the token layer reads from the same table, are left out, and a line of the dependent nouns, which have a table of their own, is an error
package parse

import "github.com/geul-org/nlgeul/pkg/data"

func loadClosed(rows []data.Row) (map[string][]string, error) {
	sets := map[string]int{helperSet: 0, determinerSet: 0, conjunctionSet: 0}
	out := map[string][]string{}
	for _, row := range rows {
		set := row.Cells[0]
		if set == boundSet {
			return nil, row.Errf(0, "의존명사는 의존 명사 표(%s.tsv)에 적는다", boundFile)
		}
		if _, ok := sets[set]; ok {
			out[set] = append(out[set], row.Cells[1])
		}
	}
	return out, nil
}
