//ff:type feature=token type=model
//ff:what Bundle — what is sent up for one stretch of a sentence: which unit of the sentence it is in, its letters, and its branches, best first; one branch is a settled reading, more than one a contest (one of them holds)
package token

// The stretches of a unit come in order, and last the dots that close
// the unit, as a bundle of their own. A unit that holds nothing to read
// (signs only) is one bundle with its letters as they are.
type Bundle struct {
	Unit     int
	Text     string
	Branches []Branch
}
