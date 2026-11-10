//ff:func feature=rule type=helper control=sequence
//ff:what Applies — a yes/no rule's value: +1 applies fully, -1 does not apply
package rule

// The tables give no degree for "this piece is read", so a rule that
// only checks a fact returns one of the two ends.
func Applies(yes bool) float64 {
	if yes {
		return 1
	}
	return -1
}
