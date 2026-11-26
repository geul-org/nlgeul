//ff:type feature=parse type=model
//ff:what spoken — one row of the closed table of markers written as they are spoken: the letters it is written with (its form and what it puts under or fuses into the syllable before it, if anything), the name of the standard marker it stands for, and whether it is read only after the copula stem
package parse

// Form and Lead are those of an entry (a form of whole syllables has
// Lead 0, a form that begins with a fused 어/아 rule.LeadVowel). Name
// is the name the index gives the standard marker. Copula tells that
// the written form is read only after the copula stem 이 or 아니,
// though the standard marker is read after other stems too.
type spoken struct {
	Form   string
	Lead   int
	Name   string
	Copula bool
}
