//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what hyphenLast — the names of an atom with 붙임 (runs that a hyphen joins, nameSigns) put after the others: a hyphen joins words inside a name, and a dot that the name ends with names the whole (geul-goal.md is a file name with a hyphen inside it), so the dot's reading comes first; the others keep their order
package parse

import "strings"

func hyphenLast(names []string) []string {
	hyphen, _, _ := strings.Cut(nameSigns['-'], ":")
	var out, last []string
	for _, n := range names {
		if n == hyphen {
			last = append(last, n)
			continue
		}
		out = append(out, n)
	}
	return append(out, last...)
}
