//ff:func feature=rule type=model control=sequence
//ff:what Validate — a shape must ask for two runs or more
package rule

import "errors"

func (s *Shape) Validate() error {
	if s.Min < 2 {
		return errors.New("rule: shape of fewer than two runs")
	}
	return nil
}
