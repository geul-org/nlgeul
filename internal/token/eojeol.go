//ff:type feature=token type=model
//ff:what Eojeol — one word of a reading as it is sent up: its parts (the stem first, then the markers after it), and where it ends in the unit (a token offset)
package token

type Eojeol struct {
	Parts []Part
	End   int
}
