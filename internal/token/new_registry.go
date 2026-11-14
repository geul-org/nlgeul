//ff:func feature=token type=helper control=sequence
//ff:what newRegistry — an empty registry whose numbers of birth begin at next (a scan that goes on where another stopped, as the letters after a sign of a pair are read after the unit as a whole, begins at the other's next)
package token

func newRegistry(next int) *registry {
	return &registry{
		born:   map[wordKey]int{},
		next:   next,
		states: map[stateKey]*path{},
		root:   &path{verdict: 1},
		lone:   map[wordKey]bool{},
	}
}
