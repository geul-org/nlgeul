package token

import (
	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

// copulaClosings are the closings the copula stands right after: the
// linking 어서 (아서, 여서 and the fused one: 늦어서이다, 위해서다,
// 가서일까), whose clause the copula then states as a reason; the
// dictionary's 이다 gives this use after -어서. The copula's stem may
// be dropped there as it is after a vowel (위해서+(이)+다).
var copulaClosings = closingNames(parse.Marker{Form: "서", Lead: rule.LeadVowel})
