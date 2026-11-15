//ff:func feature=token type=model control=sequence
//ff:what caseRows.Validate — the criterion must hold at least one row
package token

import "errors"

func (c *caseRows) Validate() error {
	if len(c.Rows) == 0 {
		return errors.New("token: case rows without a row")
	}
	return nil
}
