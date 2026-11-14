//ff:type feature=token type=model
//ff:what spacingRule — one spacing rule: its name (the number of its article of the spelling rules) and whether a boundary is a place of that article
package token

type spacingRule struct {
	name string
	fits func(gap) bool
}
