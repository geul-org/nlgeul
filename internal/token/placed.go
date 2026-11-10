//ff:func feature=token type=helper control=sequence
//ff:what placed — may a piece of one place stand before the piece of another place to its right: the places go up, and greater numbers may follow one another (3천만)
package token

func placed(place, next int) bool {
	return next == 0 || place < next || (place == next && place == placeBig)
}
