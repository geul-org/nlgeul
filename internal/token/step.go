//ff:type feature=token type=model
//ff:what step — one word of a reading as the search keeps it: where it ends, the word itself (its parts), how it is printed (kept for telling readings apart), its verdict (the least of its pieces') and its preference (their sum), whether it ends in a modifying closing, whether it is a determiner, whether it ends in the linking closing 어/아, and how much of its cost is given back when it follows such a word (an auxiliary verb written onto its main verb)
package token

type step struct {
	end     int
	word    Eojeol
	sig     string
	verdict float64
	pref    float64
	adn     bool
	det     bool
	conn    bool
	back    float64
}
