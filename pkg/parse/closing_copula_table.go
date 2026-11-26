package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
)

// copulaClosings are the closings the copula stands right after: the
// linking 어서 (아서, 여서 and the fused one: 늦어서이다, 위해서다,
// 가서일까), whose clause the copula then states as a reason; the
// dictionary's 이다 gives this use after -어서. The copula's stem may
// be dropped there as it is after a vowel (위해서+(이)+다).
var copulaClosings = closingNames(Marker{Form: "서", Lead: rule.LeadVowel})

// copulaNamings are the naming closings the copula is read right after
// (없+음+이+라, 가리키+(ㅁ)+이+(ㄴ)): 음 and ㅁ, after which a particle
// chain is read as well (chainHeads). The naming closing 기 is not
// among them: the words it makes before the copula are mostly nominals
// of the word table (잡기다, 오기임을), and those rows read them.
var copulaNamings = []string{namingPrefix + "음", nounName}
