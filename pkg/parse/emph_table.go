package parse

// emphObjName is 를 of emphasis right after an adverbial particle
// (집에를 가다, 어디로를 가니): the auxiliary particle 를 Urimalsaem gives
// after the particles 에 and 으로. It stands where the emphatic 가
// stands (rule.LinkEmph: after an adverbial or a directional particle,
// with nothing after it), and only after the particles of emphObjAfter
// (emphKept): not after 서 or 에서 (지정서를 is a noun with 를). The
// tables give it no count.
const emphObjName = "보조사:를"

var emphObjAfter = []string{"부사격:에", "부사격:으로", "부사격:로"}
