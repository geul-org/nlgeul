//ff:func feature=token type=helper control=sequence
//ff:what cloneNoWindow — a copy of an index that works without the window (the tables are shared)
package token

func cloneNoWindow(ix *index) *index {
	c := *ix
	c.noWindow = true
	return &c
}
