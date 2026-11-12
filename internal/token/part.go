//ff:type feature=token type=model
//ff:what Part — one part of a word as it is sent up: its letters, what it is written out with that the letters do not show, its readings, whether it is the word's stem (else a marker: a particle, a slot, a closing, a unit, a sign), whether it leaves no trace in the letters at all, and for a stem written wrong the letters it is written with
package token

// Text of a stem is the stem as it is looked up: with its last syllable
// as it was before a piece was fused into it, or the stem of the word
// table a changed form stands for, or the right form of a stem that is
// written wrong; Written is then the stem as it is written (else ""). Hidden is "ㄴ" for (ㄴ)다 and "이"
// for a dropped (이). A stem no table knows has no names.
type Part struct {
	Text    string
	Hidden  string
	Names   []Name
	Stem    bool
	Unseen  bool
	Written string
}
