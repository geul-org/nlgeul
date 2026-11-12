//ff:func feature=rule type=model control=sequence
//ff:what SpecName — the criterion's name: "list:" and the list's name
package rule

func (l *List) SpecName() string {
	return "list:" + l.Name
}
