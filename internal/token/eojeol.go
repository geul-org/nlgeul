//ff:type feature=token type=model
//ff:what Eojeol — one word of a reading as it is sent up: its parts (the stem first, then the markers after it), where it ends in the unit (a token offset), and whether it is the reading of a whole unit as markers alone, written apart from the word they belong to (로, 에서, 인데 after a space): such a word has no stem and attaches to the unit before
package token

type Eojeol struct {
	Parts  []Part
	End    int
	Attach bool
}
