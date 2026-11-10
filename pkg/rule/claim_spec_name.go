//ff:func feature=rule type=model control=sequence
//ff:what SpecName — the criterion's name in a rule's identifier: "claim"
package rule

func (c *Claim) SpecName() string {
	return "claim"
}
