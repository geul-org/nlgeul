//ff:func feature=mini type=helper control=sequence
//ff:what applies — a yes/no rule's value: +1 applies fully, -1 does not apply
package mini

// The tables give no degree for "this piece is read", so a rule that
// only checks a fact returns one of the two ends.
func applies(yes bool) float64 {
	if yes {
		return 1
	}
	return -1
}
