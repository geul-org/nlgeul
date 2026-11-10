//ff:type feature=rule type=model
//ff:what Word — the facts read from one word, which the rules judge: its text, the particle chain, every tail reading, the longest tail, whether it is the last word, whether it or the next word begins with a dependent noun, and the sentence
package rule

type Word struct {
	Text      string
	Chain     Chain
	Tails     []Tail
	Longest   int
	Last      bool
	Bound     bool
	NextBound bool
	Sent      *Sentence
}
