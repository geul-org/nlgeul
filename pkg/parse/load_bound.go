//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loadBound — the dependent nouns of the closed set read from the lines of the dependent-noun table (word, name, whether it is read when written onto a modifying closing too, and two columns loadBoundNouns reads): under boundSet the words that bear the set's name, under joinedMark those of them so marked, in the order of their lines; a name other than the set's or -, and a mark other than 붙여 읽음 or on a line with no name, are errors
package parse

import "github.com/geul-org/nlgeul/pkg/data"

func loadBound(rows []data.Row) (map[string][]string, error) {
	out := map[string][]string{}
	for _, row := range rows {
		word, name, mark := row.Cells[0], row.Opt(1), row.Opt(2)
		if name != "" && name != boundSet {
			return nil, row.Errf(1, "%s 또는 -를 적는다", boundSet)
		}
		if mark != "" && (mark != joinedMark || name == "") {
			return nil, row.Errf(2, "이름표가 %s인 줄에만 \"%s\"을 적고, 다른 줄에는 -를 적는다", boundSet, joinedMark)
		}
		if name != "" {
			out[boundSet] = append(out[boundSet], word)
		}
		if mark != "" {
			out[joinedMark] = append(out[joinedMark], word)
		}
	}
	return out, nil
}
