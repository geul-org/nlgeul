package token

import (
	"github.com/geul-org/nlgeul/data"
)

// closingCounts are the observed counts of the closings, form by form,
// that only this layer reads; nlgeul parse keeps the counts of its own
// table. A closing that is both a final and a linking one has the two
// counts summed. A form with no entry has no count, and a count of 0
// is taken as none (pref). An allomorph has its own count where one
// was observed apart; some pairs share one number (세요 and 으세요,
// 십시오 and 으십시오, the three forms of 어다가, ㄴ들 and 은들,
// ㅂ시다 and 읍시다, ㄹ뿐더러 and 을뿐더러). Nothing in this
// repository records where the counts were taken from or when.
var closingCounts = data.Must(data.Counts("closing-counts"))

// particleCounts are the observed counts of the particles and the
// suffixes read with them, form by form: the two forms of a pair (이 and
// 가, 을 and 를) have each its own count, but for the pairs that share
// one number (이라고 and 라고, 아 and 야, 이시여 and 시여, 이든가 and
// 든가). Nothing in this repository records where the counts were taken
// from or when.
var particleCounts = data.Must(data.Counts("particle-counts"))
