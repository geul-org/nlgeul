//ff:type feature=token type=model
//ff:what Options — how a reader is set up: the path of a word table (none when empty) how many layers of it are used (0 means all), and whether the readings that are asleep are printed too
package token

type Options struct {
	Words  string
	Layers int
	Asleep bool
}
