package token

// The shape of a court case number (사건번호): the year the case was
// filed (caseYearLens: two digits, or four), the mark of the kind of
// case (1 to caseMarkMax Hangul syllables: 다, 도, 회합, 카기전, as the
// Supreme Court's rule on case marks gives them), and the serial
// number, written with no space between (2004다25123). caseMarkMax is the
// most syllables a mark of that rule has.
var caseYearLens = []int{2, 4}

const caseMarkMax = 3

// branchJoint is the particle that joins a branch number to the number it
// branches from (가지번호: 제5조의2, 별표 11의2). Between two runs of
// digits it is that particle, not a case mark.
const branchJoint = "의"
