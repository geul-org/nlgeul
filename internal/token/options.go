//ff:type feature=token type=model
//ff:what Options — how a reader is set up: the path of a word table (none when empty), how many layers of it are used (0 means all), and the spacing rules that are on (names of SpacingArticles; none when empty)
package token

type Options struct {
	Words  string
	Layers int
	// Spacing are the spacing rules that are on, by the number of their
	// article of the spelling rules. A reader with none on reads as it
	// did before the rules.
	Spacing []string
}
