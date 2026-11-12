//ff:func feature=token type=reader control=iteration dimension=1
//ff:what ownNames — the names of a closed-set word as the sets it stands as: of the names it has (one for each set it is in), those whose set's kind is among the kinds that took what follows; none when it stands as none of its sets (then the word table's kinds name it)
package token

import "strings"

func ownNames(c *entry, passed int) []string {
	if c == nil {
		return nil
	}
	var out []string
	for _, name := range strings.Split(c.name, "|") {
		if closedKinds[labelOf(name)]&passed != 0 {
			out = append(out, name)
		}
	}
	return out
}
