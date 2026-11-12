//ff:func feature=token type=reader control=sequence
//ff:what typoClosingEntry — the entry that reads a closing written wrong, made from the entry of the right closing: a copy of it with the form and lead it is written with, the letters it is written with, and the right closing as it is printed
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// Everything but the form, the lead, the place it attaches after, wrote
// and right is the right closing's own (its name, class, count, head and
// the ways of the copula), so that what is asked by name holds for it as
// for the right one. A form under a final consonant is read after any
// syllable that has it; a form of whole syllables after what the right
// closing is read after. right is written in brackets when the right
// closing is itself a final consonant under the syllable before it.
func typoClosingEntry(right *entry, row typoClosing) *entry {
	e := *right
	e.form, e.lead = row.Form, row.Lead
	if row.Lead != 0 {
		e.after = rule.AfterAny
	}
	e.wrote = rule.LeadText(row.Lead) + row.Form
	e.right = right.form
	if right.lead != 0 {
		e.right = "(" + rule.LeadText(right.lead) + ")" + right.form
	}
	return &e
}
