package token

import "github.com/geul-org/nlgeul/pkg/parse"

// naStem is the stem of 나다, which the dictionary gives as an auxiliary
// verb after -고 (-고 나다: 듣고 나니, 먹고 나서); the word tables have it
// under the verb's kind only (mainKindOf), as the answer sheet writes it.
// naLinks is that closing -고 (naLinkedRule). naTopic is the modifying 는,
// whose letters with 나 are the pronoun 나 with the topic particle 는,
// and naNamings are the naming closings -ㅁ and -기, with which 나 has the
// letters of 남다 and of the noun 남; the rule leaves those candidates
// out.
const naStem = "나"

var (
	naLinks   = closingNames(parse.Marker{Form: "고"})
	naTopic   = []string{"꾸밈:는"}
	naNamings = []string{"이름:ㅁ", "이름:기"}
)
