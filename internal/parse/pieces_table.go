package parse

import (
	"github.com/geul-org/nlgeul/data"
)

// closingPieces tells how a closing of the table is printed when it is
// more than one piece (endings.md sections 1 and 2); a closing not
// listed is one piece. The present 는/ㄴ/느, the polite ㅂ/습/읍, the
// recalling 더 and the quoting 고 are pieces of their own, and so is
// the modifying 는 after a quotation's 다/라. Grammatical pieces fused
// into one syllable are written apart when each is a whole syllable
// (세요 = 시 + 어 + 요, 죠 = 지 + 요).
//
// For a closing that puts a consonant under the stem's last syllable,
// the value is what follows that syllable: "+다" starts a new piece
// (간+다, 합+니다); a closing not listed stays in one piece with the
// syllable (할까, 큰데).
var closingPieces = data.Must(data.Map(piecesFile))

// quoteClosings are the closings a quotation ends with before 하다
// (간다고 하다): when the text before 하 reads as one of them, it is
// divided too.
var quoteClosings = []string{"다고", "라고", "자고", "냐고", "ㄴ다고", "는다고"}
