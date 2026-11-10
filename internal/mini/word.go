//ff:type feature=mini type=model
//ff:what word — the facts read from one word, which the rules judge: its text, the particle chain, every tail reading, the longest tail, whether it is the last word, and the sentence
package mini

type word struct {
	text    string
	chain   chainFact
	tails   []tailFact
	longest int
	last    bool
	sent    *sentence
}
