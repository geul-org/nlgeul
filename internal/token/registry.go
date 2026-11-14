//ff:type feature=token type=model
//ff:what registry — what a scan keeps of the words it makes: the number each word was born with (by where it stands and how it is printed) and the next to give, for the net the state of each kind of word (stateKey) in the order they were made and the root they all grow from, the words that stand where another word printed the same stands with other values (twins, counted once for each word more than the first)
package token

type registry struct {
	born   map[wordKey]int
	next   int
	states map[stateKey]*path
	all    []*path
	root   *path
	lone   map[wordKey]bool
	twins  int
}
