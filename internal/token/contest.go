//ff:type feature=token type=model
//ff:what Contest — what is sent up for one stretch of a sentence: which unit of the sentence it is in, whether no space stood before it (Glued, not printed), its letters, and its branches, best first; one branch is a settled candidate, more than one are candidates to choose from (one of them holds)
package token

// The stretches of a unit come in order, and last the dots that close
// the unit, as a contest of their own. A unit that holds nothing to read
// (signs only) is one contest with its letters as they are.
//
// Glued tells that no space stands between this contest and the one before
// it in the sentence (false for the first); it is not printed.
type Contest struct {
	Unit     int
	Glued    bool
	Text     string
	Branches []Branch
}
