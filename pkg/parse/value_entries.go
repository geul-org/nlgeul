//ff:func feature=parse type=rule control=sequence
//ff:what valueEntries — the entries a sign of the tables of special letters is read by through the letter it writes (값): the named candidates of the sign table for that letter (？ writes ?: 물음표; ． writes a full stop: 마침표, 줄임표), when it writes one letter and the sign table does not name the sign itself; none otherwise (a unit is not looked for by what a sign writes: ㈜ writes 주 and is no week)
package parse

func (ix *index) valueEntries(f form) []*entry {
	v := []rune(f.value)
	if len(v) != 1 || len(ix.marks[f.letter]) > 0 {
		return nil
	}
	return namedEntries(ix.marks[v[0]])
}
