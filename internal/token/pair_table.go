package token

// Signs that open and close a pair (written with no space inside:
// opening signs onto the word after, closing signs onto the word
// before). The two of a pair are matched over the whole sentence, so
// that what stands between them is read on its own and a particle
// written onto the closing sign is read as one (함수(인자)를). A span
// of code between backticks is one piece and is not read. quoteSigns
// are paired two by two; round brackets by nesting.
var quoteSigns = []rune{'"', '\''}
var pairNames = map[rune]string{'(': "묶음표", ')': "묶음표", '"': "큰따옴표", '\'': "작은따옴표"}

const codeName = "코드"
