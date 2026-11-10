//ff:func feature=token type=helper control=sequence
//ff:what Asleep — does a bundle hold a unit's readings that are asleep (kept, not sent up to choose from)
package token

func (b Bundle) Asleep() bool {
	return len(b.Branches) > 0 && b.Branches[0].Asleep
}
