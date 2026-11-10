//ff:func feature=token type=reader control=sequence
//ff:what auxStems — the stems of the auxiliary verbs read right after a main verb's 어/아 form: those of the closed set (지 apart), and with a word table 가, 내 and 드리 too
package token

func (ix *index) auxStems() []string {
	if ix.lexed {
		return append(append([]string{}, auxAfter...), auxMore...)
	}
	return auxAfter
}
