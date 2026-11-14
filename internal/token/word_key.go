//ff:type feature=token type=model
//ff:what wordKey — a word of a unit by where it begins and ends and how it is printed: what is the same word whatever the words before it
package token

type wordKey struct {
	start int
	end   int
	sig   string
}
