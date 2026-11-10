//ff:type feature=token type=model
//ff:what Options — how a reader is set up: the path of a word table (none when empty) and how many layers of it are used (0 means all)
package token

type Options struct {
	Words  string
	Layers int
}
