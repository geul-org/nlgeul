package mini

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
var closingPieces = map[string]string{
	"다고": "다+고", "라고": "라+고", "자고": "자+고", "냐고": "냐+고",
	"다는": "다+는", "라는": "라+는",
	"는다": "는+다", "는다고": "는+다+고", "는다면": "는+다면", "는다는": "는+다+는",
	"는구나": "는+구나", "는군": "는+군", "느냐": "느+냐",
	"더라": "더+라", "더니": "더+니", "더라도": "더+라도",
	"습니다": "습+니다", "습니까": "습+니까", "읍시다": "읍+시다",
	"세요": "시+어+요", "시오": "시+오", "죠": "지+요",
	"ㄴ다": "+다", "ㄴ다고": "+다+고", "ㄴ다면": "+다면", "ㄴ다는": "+다+는",
	"ㅂ니다": "+니다", "ㅂ니까": "+니까", "ㅂ시다": "+시다",
}

// quoteClosings are the closings a quotation ends with before 하다
// (간다고 하다): when the text before 하 reads as one of them, it is
// divided too.
var quoteClosings = []string{"다고", "라고", "자고", "냐고", "ㄴ다고", "는다고"}
