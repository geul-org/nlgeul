package token

import (
	"github.com/geul-org/nlgeul/data"
)

// The candidates of the full stop that join the runs on its two sides
// into one name, and those that end what stands before it. Each is a
// graph of its own (dotGraph). A joining candidate's graph is named with
// the sign (파일명:.); a closing candidate is named as it is printed, with
// no sign after it (마침표).
var joinNames = []string{"파일명", "도메인", "소수점", "버전"}
var stopNames = []string{"마침표", "줄임표"}

// extensions are the file-name extensions the candidate 파일명 knows, and
// topDomains the top-level domains the candidate 도메인 knows. They are
// closed tables of common ones, not complete.
var extensions = data.Must(data.Column("extensions", 1))
var topDomains = data.Must(data.Column("top-domains", 1))

// linkSigns are signs that bind two words written with no space into
// one stem: a slash sets two things side by side (입력/출력), a hyphen
// joins two words that belong together (원-달러). Each stands only
// with a letter on both sides, and the stem only when every word in it
// is known (bound). Between two runs of Latin letters and digits the
// same signs join the runs into one name before the unit is read
// (nameSigns).
var linkSigns = map[rune]string{'/': "대비:/", '-': "붙임:-"}
var linkOrder = []rune{'/', '-'}

// run names the graph of a run of Latin letters and digits that no
// table lists. forceRun is the force of the rebuttal "it is in no key
// table" against it: lower than a new Hangul word's, because the
// change of script at its edges is itself a sign of a boundary. It is
// a placeholder, not a measured value.
const (
	runName  = "열:영숫자"
	forceRun = 0.25
)
