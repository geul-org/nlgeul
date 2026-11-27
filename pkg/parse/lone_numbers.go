//ff:func feature=parse type=reader control=sequence
//ff:what loneNumbers — the candidate of a stem whose letters are a number that is no run of digits (writtenNumber: 사십, 오, 수십, 十) read as a number with nothing written onto it: named as writtenNumber names it, with the verdict of the graph loneNumberGraph, a new word's and the kind that stands behind the words of the word table (kindBehind: before a unit written apart the letters may be its number, 사십 년); a stem with a piece on it is not read so, and numerals written in Hangul are read so only with a word table (with none, nothing tells them from the syllables of any word)
package parse

func (ix *index) loneNumbers(stem string, t tail) []kindCandidate {
	name := writtenNumber(stem)
	if name == "" || len(t.pieces) > 0 || (!ix.lexed && sinoTokens(stem) > 0) {
		return nil
	}
	return []kindCandidate{{kindBehind, name, ix.loneNumber}}
}
