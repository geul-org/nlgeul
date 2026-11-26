package parse

// The shape of a court case number (사건번호): the year the case was
// filed (caseYearLens: two digits, or four), the mark of the kind of
// case (1 to caseMarkMax Hangul syllables: 다, 도, 회합, 카기전, as the
// Supreme Court's rule on case marks gives them), and the serial
// number, written with no space between (2004다25123). caseMarkMax is the
// most syllables a mark of that rule has.
var caseYearLens = []int{2, 4}

const caseMarkMax = 3

// caseKind names a court case number. While a unit is read the whole
// number is one stem named so (an atom); it is sent up as three pieces
// of one word (caseNumberParts): the year and the serial number, each a
// number (numName), and between them the mark, named caseKind with the
// mark as its form (2005[수]+두[사건:두]+1329[수]), as a unit is named
// with its letters (개[단위:개]).
const caseKind = "사건"

// branchJoint is the particle that joins a branch number to the number it
// branches from (가지번호: 제5조의2, 별표 11의2). Between two runs of
// digits it is that particle, not a case mark.
const branchJoint = "의"
