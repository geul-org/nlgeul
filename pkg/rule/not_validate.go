//ff:func feature=rule type=model control=sequence
//ff:what Validate — the criterion has nothing to check
package rule

func (n *Not) Validate() error {
	return nil
}
