package parse

import "github.com/geul-org/nlgeul/pkg/data"

// The tables of special letters: the letters that write other letters
// in a shape of their own (① is 1 in a circle, ㈎ is 가 in round
// brackets, Ⅲ is 3 as a Roman numeral). A line gives a letter its shape
// (꼴), what it writes (내용), the letters it writes (값) and the number
// it stands for; the shape and what it writes are the letter's name in
// the output as they stand (①[원:수]). special-forms.tsv is made from
// the Unicode Character Database by the program in pkg/data/gen and
// is not written by hand; special-etc.tsv holds the letters that table
// lacks, written by hand in the same columns (its lines come first).
// Nothing here tells how many shapes, contents or letters there are:
// the tables do.
var formList = withDefaults(data.Must(loadForms(data.Must(data.Rows("special-etc", 5)), data.Must(data.Rows("special-forms", 5)))), data.Must(loadForms(data.Must(data.Rows(defaultTable, 5)))))
var forms = formsByLetter(formList)

// romanShape is the shape (꼴) the table gives the Roman numeral letters
// of Unicode (Ⅰ-Ⅻ, Ⅼ, Ⅽ, Ⅾ, Ⅿ, the small ⅰ-ⅿ and the old ↀ-ↈ): each is
// a Roman numeral by itself, and a run of them alone is a number named
// 수:로마십진, as a run of digits is a number (romanLetters).
const romanShape = "로마 숫자"

// formNumberNames are the names of the letters of the tables that write
// a number (the content 수: 원:수, 괄호:수, 로마 숫자:수 …): a stem of
// such a name is a number (numberName).
var formNumberNames = contentNames(formList, numName)

// loneFormGraph names the graph of a letter of the tables that is a
// letter of a run of Latin letters and digits (a Roman numeral letter,
// a full-width letter or digit) read alone by its shape (loneForms):
// the run's own reading stands before it.
const loneFormGraph = "꼴 글자(홀로 섬)"

// defaultTable is the table of the kinds signs fall into when no other
// table names them (special-groups.tsv, made from the Unicode Character
// Database as special-forms.tsv is): an arrow with its direction
// (→[화살표:오른쪽]), a box-drawing line (┌[기호:상자 그림]), an emoji
// (😅[기호:이모지]), a shape, a sign of mathematics, a sign of money, and
// any other sign (기호:그 밖). A line of it is read only for a letter
// that special-etc.tsv and special-forms.tsv do not hold (withDefaults)
// and that no hand-written table names (addForms). madeTable is the
// table of the shapes (special-forms.tsv): a sign it names is a named
// sign (namedSign); a sign named by the other two tables is read as a
// sign with no name was (signRun: ○○ stands for a name kept back).
const (
	defaultTable = "special-groups"
	madeTable    = "special-forms"
)

// hideKind and bulletKind are the kinds (the first column of a name)
// whose candidates are judged by where the sign stands (formGraph):
// 가림, a sign written for what is kept back (○○○, △△동, ××시), and
// 글머리, a shape that marks an item at the head of a line (● 첫째).
// forceNoHide is the force of the rebuttal that lowers 가림 where
// nothing in the sign's place tells it (one sign alone); it is the
// force of the other rebuttals of a sign (forceDashSum) and was not
// fitted to the answer sheet.
const (
	hideKind    = "가림"
	bulletKind  = "글머리"
	forceNoHide = forceDashSum
)
