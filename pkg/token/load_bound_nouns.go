//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadBoundNouns — the lists this layer reads from the lines of the dependent-noun table (word, and the two columns pkg/parse reads, then where the word stands and what else reads it): under each name of a place (boundPlaces) and of a use (boundUses) the words written with it, in the order of their lines; a name of neither set is an error, and so is a line with no place
package token

import "github.com/geul-org/nlgeul/pkg/data"

func loadBoundNouns(rows []data.Row) (map[string][]string, error) {
	out := map[string][]string{}
	for _, row := range rows {
		places, err := row.Flags(3, boundPlaces)
		if err != nil {
			return nil, err
		}
		if places == 0 {
			return nil, row.Errf(3, "자리를 하나 이상 적는다")
		}
		uses, err := row.Flags(4, boundUses)
		if err != nil {
			return nil, err
		}
		addFlagged(out, row.Cells[0], places, boundPlaces)
		addFlagged(out, row.Cells[0], uses, boundUses)
	}
	return out, nil
}
