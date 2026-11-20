//ff:func feature=parse type=reader control=sequence
//ff:what splitsOf — the pieces a closing of the tables is written out as (splits), and with a word table also those of 시오 and 으시오, the honorific and the 하오체 오 (haoSplits), whose graph 맺음:오 the index has only then; nil for a closing written whole
package parse

func (ix *index) splitsOf(name string) []split {
	if parts, ok := haoSplits[name]; ok && ix.lexed {
		return parts
	}
	return splits[name]
}
