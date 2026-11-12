//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loadClosed — the closed sets of words read from the lines of their table (set, word, and for a dependent noun whether it is read when written onto a modifying closing too): the words of each set in the order of their lines, and under joinedMark the dependent nouns so marked
package parse

import "github.com/geul-org/nlgeul/data"

func loadClosed(rows []data.Row) (map[string][]string, error) {
	sets := map[string]int{boundSet: 0, helperSet: 0, determinerSet: 0, conjunctionSet: 0}
	out := map[string][]string{}
	for _, row := range rows {
		if _, err := row.Pick(0, sets); err != nil {
			return nil, err
		}
		set, word, mark := row.Cells[0], row.Cells[1], row.Opt(2)
		if mark != "" && (mark != joinedMark || set != boundSet) {
			return nil, row.Errf(2, "%s의 줄에만 \"%s\"을 적고, 다른 줄에는 -를 적는다", boundSet, joinedMark)
		}
		out[set] = append(out[set], word)
		if mark != "" {
			out[joinedMark] = append(out[joinedMark], word)
		}
	}
	return out, nil
}
