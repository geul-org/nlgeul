//ff:func feature=token type=rule control=iteration dimension=1
//ff:what whyOf — the reason an exception carries on its node (the name of its Claim); "" when it carries none
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func whyOf(specs toulmin.Specs) string {
	for _, s := range specs {
		if c, ok := s.(*rule.Claim); ok {
			return c.Name
		}
	}
	return ""
}
