package token

// Letters that hold two grammar pieces (decision 014-29, the simple
// ones: one reading, rare inside words, the stem ends before them).
//
// honorLetters hold the honorific 시 and what is fused under it: 셨
// (시+었), 셔 (시+어…), 십 (시+ㅂ…). Wherever such a letter is read as a
// stem's last syllable 시 with a piece fused into it, it is read as
// the honorific too. After a consonant the link vowel 으 stands before
// it (읽으셨다).
const (
	honorName  = "선어말:시"
	honorCount = 19566
	euName     = "이음:으"
	negName    = "합침:잖"
	negLetter  = '잖'
	jiName     = "맺음:지"
)

var honorLetters = []rune{'셨', '셔', '십'}

// negParts is what 잖 is written out as: the closing 지 and the stem
// of 않다, which the closing after it attaches to.
var negParts = []part{{"지", jiName}, {"(않)", "부정:않"}}

// splits are closings of the tables that are written out as their
// pieces: 죠 is 지+요, 십시오 the honorific and ㅂ시오.
var splits = map[string][]part{
	"맺음:죠":    {{"지", jiName}, {"(요)", "보조사:요"}},
	"맺음:십시오":  {{"시", honorName}, {"(ㅂ)시오", "맺음:ㅂ시오"}},
	"맺음:으십시오": {{"으", euName}, {"시", honorName}, {"(ㅂ)시오", "맺음:ㅂ시오"}},
}

// contracts are the fused forms whose first word is fixed, a closed
// set: only what is listed is read. The last three are the dependent
// noun 거 with a particle under it, written onto a modifying closing
// (한건, 하는게, 할걸): they are read only right after such a closing,
// next to the closings and words of the same letters.
var contracts = []contract{
	{"이게", "이것", "(이)", "주격:이", '이', false}, {"그게", "그것", "(이)", "주격:이", '이', false}, {"저게", "저것", "(이)", "주격:이", '이', false},
	{"이건", "이것", "(은)", "보조사:은", '은', false}, {"그건", "그것", "(은)", "보조사:은", '은', false}, {"저건", "저것", "(은)", "보조사:은", '은', false},
	{"이걸", "이것", "(을)", "목적격:을", '을', false}, {"그걸", "그것", "(을)", "목적격:을", '을', false}, {"저걸", "저것", "(을)", "목적격:을", '을', false},
	{"뭘", "뭐", "(를)", "목적격:를", '를', false}, {"무얼", "무엇", "(을)", "목적격:을", '을', false},
	{"누가", "누구", "가", "주격:가", '가', false}, {"누굴", "누구", "(를)", "목적격:를", '를', false},
	{"여긴", "여기", "(는)", "보조사:는", '는', false}, {"여길", "여기", "(를)", "목적격:를", '를', false},
	{"우린", "우리", "(는)", "보조사:는", '는', false}, {"우릴", "우리", "(를)", "목적격:를", '를', false},
	{"그땐", "그때", "(는)", "보조사:는", '는', false}, {"땐", "때", "(는)", "보조사:는", '는', false},
	{"게", "거", "(이)", "주격:이", '이', true},
	{"건", "거", "(는)", "보조사:는", '는', true},
	{"걸", "거", "(를)", "목적격:를", '를', true},
}
