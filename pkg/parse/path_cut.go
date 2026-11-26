//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what pathCut — where a path that a unit begins with is cut from the Hangul right after its last slash: the path's end, when the path ends in a slash and Hangul syllables alone stand from there to the unit's end (~/a/b/회의록, ~/a/b/폴더는); 0 when the unit begins with no path, the path is a slash alone (/좌측: the sign 빗금 before a word), the path does not end in a slash, or nothing or something else follows it
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The letters do not tell whether that Hangul is a name of the path
// (a file's name) or a word of the sentence written onto a directory's
// slash, so both are read: the path up to the unit's end (namedTail)
// and the path cut here with a word after it (whole).
func pathCut(s []rune, name string, end int) int {
	if name != pathName || end < 2 || end >= len(s) || s[end-1] != '/' {
		return 0
	}
	for _, r := range s[end:] {
		if rule.JongOf(r) < 0 {
			return 0
		}
	}
	return end
}
