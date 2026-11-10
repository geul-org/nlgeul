//ff:type feature=parse type=model
//ff:what Marker — one row of the grammar tables as other packages read it: a particle or a closing with how it attaches and its observed count
package parse

// Lead is 0 for whole syllables, -1 for a closing that begins with
// 어/아/여, else the final-consonant index it puts under the syllable
// before it (4 ㄴ, 8 ㄹ, 16 ㅁ, 17 ㅂ). After is what must precede a
// whole-syllable form: 0 anything, 1 a consonant, 2 a consonant other
// than ㄹ, 3 a vowel or ㄹ, 4 a vowel. Slot is a particle's place in a
// chain; Kase its case flags; Kind a closing's kind flags (the rule
// package's constants). Eu: takes 으 after a consonant. NoYo: 요 does
// not follow. NoIda: does not attach to the copula stem 이. NeedIda:
// attaches only to 이 or 아니. SelfIda: holds the copula stem itself.
type Marker struct {
	Form    string
	Lead    int
	After   int
	Slot    int
	Kase    int
	Kind    int
	Eu      bool
	NoYo    bool
	NoIda   bool
	NeedIda bool
	SelfIda bool
	Count   int
}
