//ff:type feature=token type=model
//ff:what eojeol — one word of a reading: where it ends, how it is printed, the graph names in it, its verdict (the least of its pieces') and its preference (their sum), and whether it ends in a modifying closing
package token

type eojeol struct {
	end     int
	marked  string
	names   string
	verdict float64
	pref    float64
	adn     bool
}
