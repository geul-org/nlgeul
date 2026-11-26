//ff:func feature=parse type=reader control=sequence
//ff:what spokenEntry — the entry that reads a marker written as it is spoken, made from the entry of the standard marker: a copy of it with the written form, no count of its own, the graph the word before it must pass, and for a row read only after the copula stem that condition
package parse

// Everything but the form, the count, the gate and the copula condition
// is the standard marker's own (its name, class, lead, place, kind in a
// chain and head), so that what is asked by name holds for it as for the
// standard one, and the piece is printed with the standard name. Its
// count is 0: its preference is what its syllables would cost as stem
// syllables (pref).
func spokenEntry(standard *entry, row spoken, gate *entry) *entry {
	e := *standard
	e.form, e.count, e.gate = row.Form, 0, gate
	if row.Copula {
		e.needs = true
	}
	return &e
}
