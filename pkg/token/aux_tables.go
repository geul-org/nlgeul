package token

import "github.com/geul-org/nlgeul/pkg/parse"

// malStem is the stem of 말다, which the dictionary gives as an
// auxiliary verb after -지 (가지 마라: do not) and after -고 (지고 말았다:
// in the end); the word tables have it under the verb's kind only
// (mainKindOf), as the answer sheet writes it. malLinks are those two
// closings (malLinkedRule).
const malStem = "말"

var malLinks = closingNames(parse.Marker{Form: "지"}, parse.Marker{Form: "고"})

// regardStems are the stems of the verbs that after …에 name what is
// talked or thought about (관하다, the dictionary's "말하거나 생각하는
// 대상으로 하다": …에 관하여, …에 관한). The verb after their -여 form is
// a main verb that takes that phrase, not an auxiliary of theirs (…에
// 관하여 보건대). Only 관하다 is listed, the one the answer sheet shows.
var regardStems = []string{"관하"}
