//ff:func feature=parse type=helper control=sequence
//ff:what plainToo — the candidates of an atom's stem, with the new word of the same letters after them when the atom is two names joined by one slash and read as a path for that slash alone (I/O, src/main: no head of a path, no second slash): the letters do not tell a path from a word written with a slash, so the reading with no name stays beside the path; any other atom's candidates as they are
package parse

import "strings"

func plainToo(stem string, candidates []stemCandidate) []stemCandidate {
	if len(candidates) != 1 || candidates[0].name != pathName || strings.Count(stem, "/") != 1 || isPath([]rune(stem)) {
		return candidates
	}
	return append(candidates, stemCandidate{kindCandidate{}, stem, ""})
}
