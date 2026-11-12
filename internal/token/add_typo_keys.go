//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addTypoKeys — put the closed table's wrong ways to write a closed-set word (typoKeys) among the misspelt forms, with the kind the closed-set word has as a member of its set; a line the word table gives for the same two forms is joined to it, not added again
package token

// It is read after the word table is loaded, so that a stem written
// wrong is judged the way a closed-set word is when it is written
// right (misspelt): the name it prints is the closed-set word's.
func (ix *index) addTypoKeys() {
	for wrote, right := range typoKeys {
		ix.typos[wrote] = typoJoined(ix.typos[wrote], right, ix.closed[right])
	}
}
