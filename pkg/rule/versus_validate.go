//ff:func feature=rule type=model control=sequence
//ff:what Validate — both sides must be named
package rule

import "errors"

func (v *Versus) Validate() error {
	if v.Mine == "" || v.Rival == "" {
		return errors.New("rule: versus without a side")
	}
	return nil
}
