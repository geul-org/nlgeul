//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadCases — the rows of the case exception table read from its lines, graph by graph, in the order of their lines; a word written for one graph on two lines is an error that names the line before
package token

import "github.com/geul-org/nlgeul/pkg/data"

// One word of one graph is one row, so that a trace finds the row that
// applied by the word (caseRows.note).
func loadCases(rows []data.Row) (map[string][]caseException, error) {
	out := map[string][]caseException{}
	seen := map[[2]string]int{}
	for _, row := range rows {
		c, err := caseExceptionOf(row)
		if err != nil {
			return nil, err
		}
		key := [2]string{c.Graph, c.Word}
		if line := seen[key]; line != 0 {
			return nil, row.Errf(1, "%d번째 줄이 같은 그래프에 이 낱말을 이미 적었다", line)
		}
		seen[key] = row.Line
		out[c.Graph] = append(out[c.Graph], c)
	}
	return out, nil
}
