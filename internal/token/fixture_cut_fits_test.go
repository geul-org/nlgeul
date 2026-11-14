//ff:func feature=token type=helper control=iteration dimension=1
//ff:what cutFits — whether the spans of a way run on from the place given (at) without a gap or a word with no letters, and end at at+n
package token

func cutFits(c cut, at, n int) bool {
	next := at
	for _, sp := range c.spans {
		if sp.Start != next || sp.End <= sp.Start {
			return false
		}
		next = sp.End
	}
	return next == at+n
}
