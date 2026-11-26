//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what loneHead — how many letters at the front of a unit of signs alone are a path's head that stands with no name after it (loneHeads: ~/, ./, ../), which is read as a path as it is; 0 when the unit begins with none, or a slash or a dot goes on right after the head (.//, ../..)
package parse

import "strings"

// loneHeads are the heads of a path that are a path with no name after
// them: the home directory, this directory, the one above. A slash
// alone is not one (it is the sign 빗금).
var loneHeads = []string{"~/", "./", "../"}

func loneHead(all []rune) int {
	text := string(all)
	for _, head := range loneHeads {
		n := len([]rune(head))
		if strings.HasPrefix(text, head) && !strings.ContainsRune("/.", atOf(all, n)) {
			return n
		}
	}
	return 0
}
