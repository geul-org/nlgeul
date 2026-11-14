//ff:func feature=token type=reader control=sequence
//ff:what auxStems — the stems of the auxiliary verbs read right after a main verb's 어/아 form (addAux)
package token

func (ix *index) auxStems() []string {
	return ix.aux
}
