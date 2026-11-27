//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what withDefaults — the lines of the tables of special letters with the lines of the table of kinds (defaultTable) added after them, each only for a letter the lines before do not hold: a letter with a line of its own is not given the kind it would fall into
package parse

func withDefaults(list, defaults []form) []form {
	held := map[rune]bool{}
	for _, f := range list {
		held[f.letter] = true
	}
	for _, f := range defaults {
		if !held[f.letter] {
			list = append(list, f)
		}
	}
	return list
}
