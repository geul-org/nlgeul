//ff:type feature=rule type=model
//ff:what Bits — a criterion of flags: which set it speaks of (case, kind, word) and the flags a fact must share with it
package rule

type Bits struct {
	Of   string
	Mask int
}
