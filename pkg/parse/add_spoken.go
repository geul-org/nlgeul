//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what addSpoken — the entry of every row of the table of markers written as they are spoken, made from the entry of the standard marker the row names and put under the key of the written form: a closing is read only after a predicate's stem of the word table, a particle only after a nominal of it; its count is not added to the total, nor its form to the closing set; a row whose standard marker the index does not have, or has with another lead, makes none
package parse

import "strings"

// The standard marker is found under the last letter of its name, which
// is its form's (a fused 어/아 closing is named with (합침) after it).
// It was put in before any written form, so it is the first of its name
// there.
func (ix *index) addSpoken() {
	for _, row := range spokenMarkers {
		name := []rune(strings.TrimSuffix(row.Name, "(합침)"))
		standard := ix.named(name[len(name)-1], row.Name)
		if standard.g == nil || standard.lead != row.Lead {
			continue
		}
		gate := ix.pred
		if standard.class == classParticle {
			gate = ix.nominal
		}
		e := spokenEntry(standard, row, gate)
		ix.edges += markerGraph(e, closedEnding(ix.words, e.form), caseExceptions[e.name], ix.lexed)
		ix.graphs++
		k := entryKey(e)
		ix.byKey[k] = append(ix.byKey[k], e)
	}
}
