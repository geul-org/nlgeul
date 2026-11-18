package token

import "slices"

// outsideStem and outsideCase are the noun 밖 and the adverbial particle
// 에, whose letters together are also the particle 밖에
// (outsideNounRule). negWords are the words of negation the particle 밖에
// is followed by: the stems of 없다, 않다, 못하다, 모르다 and 아니다, and the
// negative adverbs of the closed set (안, 못, 아니).
const (
	outsideStem = "밖"
	outsideCase = "부사격:에"
)

var negWords = slices.Concat([]string{"없", "않", "못하", "모르", "아니"}, negAdverbs)
