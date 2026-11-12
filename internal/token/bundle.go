//ff:type feature=token type=model
//ff:what Bundle — what is sent up for one stretch of a sentence: which unit of the sentence it is in, whether no space stood before it (Glued, not printed), its letters, and its branches, best first; one branch is a settled candidate, more than one a contest (one of them holds)
package token

// The stretches of a unit come in order, and last the dots that close
// the unit, as a bundle of their own. A unit that holds nothing to read
// (signs only) is one bundle with its letters as they are.
//
// Glued tells that no space stands between this bundle and the one before
// it in the sentence (false for the first); it is not printed.
type Bundle struct {
	Unit     int
	Glued    bool
	Text     string
	Branches []Branch
}
