//ff:func feature=rule type=helper control=sequence
//ff:what MarksOf — the names of the markers of a word as one of its candidates reads it (Near.Marks), each apart and in order; none for a word of a stem alone
package rule

import "strings"

func MarksOf(n Near) []string {
	if n.Marks == "" {
		return nil
	}
	return strings.Split(n.Marks, "+")
}
