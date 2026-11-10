//ff:func feature=rule type=model control=sequence
//ff:what SpecName — the criterion's name: "bits:" and the set it speaks of
package rule

func (b *Bits) SpecName() string {
	return "bits:" + b.Of
}
