//ff:func feature=rule type=model control=sequence
//ff:what Validate — a criterion of flags must name at least one
package rule

import "errors"

func (b *Bits) Validate() error {
	if b.Mask == 0 {
		return errors.New("rule: bits without a flag")
	}
	return nil
}
