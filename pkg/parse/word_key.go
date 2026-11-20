//ff:type feature=parse type=model
//ff:what wordKey — a word of a unit by where it begins and ends and how it is printed: what is the same word whatever the words before it
package parse

type wordKey struct {
	start int
	end   int
	sig   string
}
