//ff:type feature=token type=model
//ff:what typoClosing — one row of the closed table of closings written wrong: the letters it is written with (its form and the final consonant it puts under the syllable before it, if any) and the name of the right closing it stands for
package token

// Form and Lead are those of an entry (a form of whole syllables has
// Lead 0). Name is the name the index gives the right closing.
type typoClosing struct {
	Form string
	Lead int
	Name string
}
