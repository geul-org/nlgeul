//ff:func feature=token type=rule control=sequence
//ff:what keyed — whether a text is a key of the closed-set words, the word table or its wrong ways to write a word: the words a part of a bound stem may be (partVerdict, partTypo)
package token

func (ix *index) keyed(text string) bool {
	_, listed := ix.lex[text]
	_, closed := ix.closed[text]
	_, set := ix.words[text]
	_, wrong := ix.typos[text]
	return listed || closed || set || wrong
}
