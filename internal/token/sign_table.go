package token

// The signs that are named where they stand at the edges of a unit or
// alone (the full stop has its own graphs: stopNames), each with its
// readings in the order they are printed. A reading is a graph of its
// own (signGraph). A sign with no reading that stands, and a sign that
// is not in the table, is sent up as its letters with no name: no
// letter of a sentence is dropped.
var signNames = map[rune][]string{
	',': {"쉼표"},
	'?': {"물음표"},
	'!': {"느낌표"},
	':': {"쌍점"},
	';': {"쌍반점"},
	'~': {"물결표"},
	'#': {headingName, "샵", "태그"},
	'@': {"골뱅이", "이름"},
}
var signOrderNamed = []rune{',', '?', '!', ':', ';', '~', '#', '@'}

// roleNames are the readings of a sign by what it does for the words
// after it (a heading, a topic, the name of what is meant) and not as
// the sign itself. When only such readings stand, the sign's letters
// with no name are sent up beside them. headingName is printed with
// the number of the signs in a row and no colon (## is 제목2).
var roleNames = []string{headingName, "태그", "이름"}

const headingName = "제목"

// arrowSigns stand right next to < or > in a run that is no bracket
// (->, =>, <=, >=).
const arrowSigns = "-="

// stopSigns are the signs that close the word before them inside a
// unit, as a piece of that word (stopAt): the full stop and the comma.
const stopSigns = ".,"
