//ff:type feature=parse type=model
//ff:what word — the facts read from one word, which the rules judge: its text, the particle chain, every tail reading, the longest tail, whether it is the last word, whether it or the next word begins with a dependent noun, and the sentence
package parse

type word struct {
	text      string
	chain     chainFact
	tails     []tailFact
	longest   int
	last      bool
	bound     bool
	nextBound bool
	sent      *sentence
}
