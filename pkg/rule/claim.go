//ff:type feature=rule type=model
//ff:what Claim — the criterion that names a claim node: its name in the graph and the conclusion it stands for; it lets one rule function serve several claims
package rule

type Claim struct {
	Name      string
	Concludes string
}
