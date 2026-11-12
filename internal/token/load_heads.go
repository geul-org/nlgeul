//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadHeads — the signs at the front of a number read from their table: a sign and its kind a line, named with the kind and the sign (부호:-); a sign written on two lines is an error
package token

import "github.com/geul-org/nlgeul/data"

func loadHeads(rows []data.Row) (map[rune]string, error) {
	out := map[rune]string{}
	for _, row := range rows {
		sign, err := row.Rune(0)
		if err != nil {
			return nil, err
		}
		if _, ok := out[sign]; ok {
			return nil, row.Errf(0, "앞 줄에 이미 적혀 있다")
		}
		out[sign] = row.Cells[1] + ":" + row.Cells[0]
	}
	return out, nil
}
