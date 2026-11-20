//ff:func feature=parse type=model control=sequence
//ff:what caseRows.Validate — the criterion must hold at least one row
package parse

import "errors"

func (c *caseRows) Validate() error {
	if len(c.Rows) == 0 {
		return errors.New("parse: case rows without a row")
	}
	return nil
}
