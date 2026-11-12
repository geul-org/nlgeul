//ff:func feature=rule type=model control=sequence
//ff:what SpecName — the criterion's name: "not"
package rule

func (n *Not) SpecName() string {
	return "not"
}
