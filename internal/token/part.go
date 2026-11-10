//ff:type feature=token type=model
//ff:what Part — one part of a word as it is sent up: its letters, what it is written out with that the letters do not show, its readings, whether it is the word's stem (else a marker: a particle, a slot, a closing, a unit, a sign), and whether it leaves no trace in the letters at all
package token

// Text of a stem is the stem as it is looked up: with its last syllable
// as it was before a piece was fused into it, or the stem of the word
// table a changed form stands for. Hidden is "ㄴ" for (ㄴ)다 and "이"
// for a dropped (이). A stem no table knows has no names. Ends is, for
// a closing (맺음), whether it closes a sentence (rule.KindFinal),
// links to what follows (rule.KindLink), or both, as the closing
// tables mark it; 0 for any other part and for a closing the tables
// do not mark. It is not printed.
type Part struct {
	Text   string
	Hidden string
	Names  []Name
	Stem   bool
	Unseen bool
	Ends   int
}
