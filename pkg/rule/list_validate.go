//ff:func feature=rule type=model control=sequence
//ff:what Validate — a list must have a name and say which fact it is asked for
package rule

import "errors"

func (l *List) Validate() error {
	if l.Name == "" || l.Of < ListForm || l.Of > ListClosing {
		return errors.New("rule: list without a name or a fact")
	}
	return nil
}
