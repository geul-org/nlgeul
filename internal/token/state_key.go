//ff:type feature=token type=model
//ff:what stateKey — a state of the net: a word by what it is (where it stands, how it is printed, its verdict, its preference, whether it ends in a modifying closing, is a determiner, ends in the linking closing, how much of its cost is given back, and the kind of its stem); the words before it do not take part, so every candidate that ends in such a word is one state
package token

type stateKey struct {
	word    wordKey
	verdict float64
	pref    float64
	back    float64
	adn     bool
	det     bool
	conn    bool
	kind    int
}
