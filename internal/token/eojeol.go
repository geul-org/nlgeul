//ff:type feature=token type=model
//ff:what eojeol — one word of a reading: where it ends, how it is printed, the graph names in it, its verdict (the least of its pieces') and its preference (their sum), whether it ends in a modifying closing, whether it ends in the linking closing 어/아, and how much of its cost is given back when it follows such a word (an auxiliary verb written onto its main verb)
package token

type eojeol struct {
	end     int
	marked  string
	names   string
	verdict float64
	pref    float64
	adn     bool
	conn    bool
	back    float64
}
