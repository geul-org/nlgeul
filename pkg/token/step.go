//ff:type feature=token type=model
//ff:what step — one word of a candidate as the search keeps it: where it ends, the word itself (its parts), how it is printed (kept for telling candidates apart), its verdict (the least of its pieces') and its preference (their sum), whether it ends in a modifying closing, whether it is a determiner, whether it ends in the linking closing 어/아, the kind of its stem (0 for none), and when it was first made in its scan (born: its number among the words by where they begin, end and how they are printed, in the order they first came; of two of the same lambda the one made first comes first)
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
	kind    int
	born    int
}
