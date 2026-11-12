//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadPairs — the signs of a pair read from their table, sign by sign; a sign written on two lines is an error
package token

import "github.com/geul-org/nlgeul/data"

func loadPairs(rows []data.Row) (map[rune]pairRow, error) {
	out := map[rune]pairRow{}
	for _, row := range rows {
		sign, pair, err := pairOf(row)
		if err != nil {
			return nil, err
		}
		if _, ok := out[sign]; ok {
			return nil, row.Errf(0, "앞 줄에 이미 적혀 있다")
		}
		out[sign] = pair
	}
	return out, nil
}
