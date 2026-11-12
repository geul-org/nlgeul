//ff:type feature=rule type=model
//ff:what List — a criterion of a list: its name, which fact is looked up in it (List…) and the items
package rule

// The items are left out of a rule's identifier: the name tells two
// lists of one graph apart.
type List struct {
	Name  string
	Of    int
	Items []string `json:"-"`
}
