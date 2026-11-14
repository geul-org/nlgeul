package parse

import (
	"github.com/geul-org/nlgeul/data"
)

// endings are the closings of the table closings.tsv in the data
// package (from endings.md sections 1, 3 and 4 and morphemes.md
// sections 2 and 3). Counts are the ones written in the table (0 where
// it gives none). The fourth field is the closing's kind as section 3
// marks it (종, 연, 꾸, 이). Of the naming forms only ㅁ and
// 음 are read (a sentence may end in them); 기 and
// the rare closings (오, 소, 세, 마, 쇼, 람, 담, 렴, 되, 곤, 건대,
// 건만, 길래, ㄴ들, ㄴ바, ㄴ즉, ㄹ망정, ㄹ지언정, ㄹ뿐더러) are left out.
var endings = data.Must(data.Each(closingsFile, 5, endingOf))
