//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadContracts — the fused forms read from their table, set by set, in the order of their lines
package token

import "github.com/geul-org/nlgeul/data"

func loadContracts(rows []data.Row) (map[string][]contract, error) {
	out := map[string][]contract{}
	for _, row := range rows {
		c, err := contractOf(row)
		if err != nil {
			return nil, err
		}
		out[row.Cells[0]] = append(out[row.Cells[0]], c)
	}
	return out, nil
}
