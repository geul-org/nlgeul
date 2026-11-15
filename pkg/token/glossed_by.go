//ff:func feature=token type=reader control=iteration dimension=1
//ff:what glossedBy — are the letters of a unit that are read (after its head) the sound of the Han characters glossed right after them: Hangul syllables alone, as many as the characters the round bracket after the unit holds (han, glossAt) or as stand right before the bracket the unit stands alone in (soundAt), with no sign trimmed between them and the bracket (after: the signs after the letters)
package token

// A Hangul word glossed so is the sound of those characters, and none
// of its syllables is a native particle, closing or slot (관치(官治) is
// no 관하+치).
func glossedBy(letters []rune, han int, after string) bool {
	if han == 0 || after != "" || len(letters) != han {
		return false
	}
	for _, r := range letters {
		if r < '가' || r > '힣' {
			return false
		}
	}
	return true
}
