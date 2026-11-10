//ff:type feature=mini type=model
//ff:what ending — one closing form of a predicate tail: syllables, how it joins the syllable before, its kind, flags, observed count
package mini

// lead is leadNone, leadA, or the final consonant the closing puts on
// the syllable before it (jongN, jongL, jongB). after applies when
// lead is leadNone.
type ending struct {
	form  string
	lead  int
	after int
	kind  int
	flags int
	count int
}
