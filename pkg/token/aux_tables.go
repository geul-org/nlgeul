package token

import "github.com/geul-org/nlgeul/pkg/parse"

// malStem is the stem of 말다, which the dictionary gives as an
// auxiliary verb after -지 (가지 마라: do not) and after -고 (지고 말았다:
// in the end); the word tables have it under the verb's kind only
// (mainKindOf), as the answer sheet writes it. malLinks are those two
// closings (malLinkedRule).
const malStem = "말"

var malLinks = closingNames(parse.Marker{Form: "지"}, parse.Marker{Form: "고"})
