package parse

// oneStem is 한 standing alone, which is three words: the determiner
// (한 사람, 한 번), 하다 with the modifying closing ㄴ (일을 한 사람),
// and the noun 한(限) (-는 한).
const oneStem = "한"

// oneNoun is the name of 한 read as a nominal of the word table, the
// noun 한(限) of the construction -는 한 (limitOneRule), and neunAdn
// the name of the modifying closing 는 that stands before it.
var (
	oneNoun = keyKind + ":" + kindLabel[kindNominal]
	neunAdn = adnPrefix + "는"
)
