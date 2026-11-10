//ff:func feature=rule type=model control=sequence
//ff:what Validate — a claim must have a name and a conclusion
package rule

import "errors"

func (c *Claim) Validate() error {
	if c.Name == "" || c.Concludes == "" {
		return errors.New("rule: claim without a name or a conclusion")
	}
	return nil
}
