//ff:func feature=parse type=reader control=sequence
//ff:what auxStems — the stems of the auxiliary verbs read right after a main verb's 어/아 form (addAux)
package parse

func (ix *index) auxStems() []string {
	return ix.aux
}
