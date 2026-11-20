//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what caseWords — the words of rows of the case exception table, in the order of the rows: the list a graph's case exception looks the word under judgment up in
package parse

func caseWords(rows []caseException) []string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.Word)
	}
	return out
}
