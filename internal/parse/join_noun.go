//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what joinNoun — read the dependent noun at s[k:] written onto a modifying closing (아는거, 할수, 있는거지): the chain it gives (the noun alone, or with its particles) and the tails on the noun
package parse

// The lengths count the modifying closing before the noun too. A noun
// with nothing after it is a chain of the case caseBare.
func joinNoun(s []rune, k int, noun string) (chainFact, []tailFact) {
	n := len([]rune(noun))
	left, size, adn := leftAdn(s[:k])
	if got, ok := suffixAt(s, k+n, noun); !ok || got != k || !adn {
		return chainFact{}, nil
	}
	head := s[k:]
	base := size + 2*n
	if len(head) == n {
		return chainFact{base, "", "", 0, caseBare, left + "+" + noun}, nil
	}
	var tails []tailFact
	for _, t := range collectTails(head, len(head), false) {
		if t.stem == noun {
			t.size, t.marked = t.size+size, left+"+"+t.marked
			tails = append(tails, t)
		}
	}
	c := readChain(head)
	if c.size != 2*(len(head)-n) {
		return chainFact{}, tails
	}
	c.size, c.marked = c.size+base, left+"+"+c.marked
	return c, tails
}
