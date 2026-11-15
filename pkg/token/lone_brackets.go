package token

// loneBrackets are the closing brackets that close a pair with none of
// their kind open: written onto the word before them (참조), 것이다}.),
// they close an opening sign that stood before the sentence, as when a
// sentence is cut out of a passage in brackets. They are judged by
// loneBracketGraph and not as the quotation marks of loneShuts are:
// nothing has to be written right after them, and right after a digit
// they are a list number (1), 2)) and close nothing.
const loneBrackets = ")}]>」』》〉｣"

// loneBracketDepth is the depth of a closing bracket with none of its
// kind open: the opening sign it closes stood before the sentence,
// outside all that the sentence opens, so it is counted the outermost.
// In the answer sheet every such bracket the judges wrote themselves is
// at depth 1. A quotation mark of loneShuts closes at depth 0 (none of
// its kind was opened), as it did before; the two are not made one.
const loneBracketDepth = 1
