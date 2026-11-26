package parse

// loneBrackets are the closing brackets that close a pair with none of
// their kind open: written onto the word before them (참조), 것이다}.),
// they close an opening sign that stood before the sentence, as when a
// sentence is cut out of a passage in brackets. They are judged by
// loneBracketGraph and not as the quotation marks of loneShuts are:
// nothing has to be written right after them. Right after a digit they
// close no pair that stood before the sentence: a round bracket there
// numbers an item of a list (listBrackets), any other is not read.
const loneBrackets = ")}]>」』》〉｣】〕）］｝"

// listBrackets are the closing brackets that number an item of a list
// when written right after a digit with none of their kind open (1),
// 2)): the round bracket and its full-width form. Such a bracket closes
// nothing, in the sentence or before it, so it is read as a closing at
// depth 0 (the answer sheet writes every one the judges wrote
// themselves so), beside its letter with no name. Any other bracket of
// loneBrackets right after a digit is not read as a closing (5>3).
const listBrackets = ")）"

// loneBracketDepth is the depth of a closing bracket with none of its
// kind open: the opening sign it closes stood before the sentence,
// outside all that the sentence opens, so it is counted the outermost.
// In the answer sheet every such bracket the judges wrote themselves is
// at depth 1. A quotation mark of loneShuts closes at depth 0 (none of
// its kind was opened), as it did before; the two are not made one,
// but for a quotation mark right after a mark that ends a sentence
// (다.”): the quotation it closes began in a sentence before, as the
// bracket's pair did, and it closes at this depth too (the answer sheet
// writes 24 such marks at depth 1 and one at 0).
const loneBracketDepth = 1

// endMarks are the marks that end a sentence (마침표, 물음표, 느낌표,
// 줄임표). A closing sign written right after one closes on the end of
// the sentence it encloses (다.』, 하오?」): the orthography puts the
// sentence's full stop inside the closing sign (「한글 맞춤법」 부록
// 문장 부호, 마침표 (1) [붙임 1]), so it stands there as it stands right
// after a letter. A full stop right after a digit is a number's (a
// list number or a date: #2."), not the end of a sentence.
const endMarks = ".?!…"
