//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadSigns — the names of the signs read from their table: a sign and one of its candidates a line, the candidates of a sign in the order of their lines
package token

import "github.com/geul-org/nlgeul/data"

func loadSigns(rows []data.Row) (map[rune][]string, error) {
	out := map[rune][]string{}
	for _, row := range rows {
		sign, err := row.Rune(0)
		if err != nil {
			return nil, err
		}
		out[sign] = append(out[sign], row.Cells[1])
	}
	return out, nil
}
