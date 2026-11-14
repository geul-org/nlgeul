//ff:type feature=token type=model
//ff:what pool — the candidates for each span gathered from the ways of cutting (a word printed the same once, the first kept, in the order of the ways), and how many words at a place were met with different values (clashes)
package token

type pool struct {
	words   map[Span][]Branch
	clashes int
}
