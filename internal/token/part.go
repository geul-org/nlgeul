//ff:type feature=token type=model
//ff:what Part — one part of a word as it is sent up: its letters, what it is written out with that the letters do not show, its reading, whether it is the word's stem (else a marker: a particle, a slot, a closing, a unit, a sign), whether it leaves no trace in the letters at all, for a stem or a marker written wrong the letters it is written with, and for a sign that opens or closes a pair or heads a title its depth
package token

// Text of a stem is the stem as it is looked up: with its last syllable
// as it was before a piece was fused into it, or the stem of the word
// table a changed form stands for, or the right form of a stem that is
// written wrong; Written is then the stem as it is written (else ""). A
// marker written wrong (a closing) is the right marker, and Written is
// then the marker as it is written (ㄹ께). Hidden is "ㄴ" for (ㄴ)다 and "이"
// for a dropped (이). A stem no table knows has no names. Depth is the
// depth of a sign of a pair (열기:2 and 닫기:2 are 2: one pair stands
// open outside it) and the level of a heading (제목2 is 2); 0 for any
// other part.
type Part struct {
	Text    string
	Hidden  string
	Name    Name
	Stem    bool
	Unseen  bool
	Written string
	Depth   int
}
