package parse

import "github.com/geul-org/nlgeul/pkg/data"

// The tables of special letters: the letters that write other letters
// in a shape of their own (① is 1 in a circle, ㈎ is 가 in round
// brackets, Ⅲ is 3 as a Roman numeral). A line gives a letter its shape
// (꼴), what it writes (내용), the letters it writes (값) and the number
// it stands for; the shape and what it writes are the letter's name in
// the output as they stand (①[원:수]). special-forms.tsv is made from
// the Unicode Character Database by pkg/data/gen/special_forms.py and
// is not written by hand; special-etc.tsv holds the letters that table
// lacks, written by hand in the same columns (its lines come first).
// Nothing here tells how many shapes, contents or letters there are:
// the tables do.
var formList = data.Must(loadForms(data.Must(data.Rows("special-etc", 5)), data.Must(data.Rows("special-forms", 5))))
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
