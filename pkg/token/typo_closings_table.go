package token

import (
	"github.com/geul-org/nlgeul/data"
)

// typoClosings are the closings that are often written wrong, each with
// the right closing it stands for. The rows are common misspellings
// written down from memory, not copied from another list or a
// dictionary, and the lines of typo-closings.tsv in the data package are
// the whole table: a closing written wrong in some other way is not
// read. A row whose Form has a Lead is the written
// form under the final consonant ㄹ (할께 for 할게: 하 + ㄹ께); a row of
// whole syllables is read after the stems the right closing is read
// after (먹을께 for 먹을게, 했읍니다 for 했습니다). The right closing is
// found by its name among the closings of the tables; every row must
// find one. Not in the table, because the right form is not a closing of
// the tables: ㄹ려는 for 려는, ㄹ래야 for 려야, ㄹ런지 and ㄹ른지 for
// ㄹ는지, 는구만 for 는구먼.
var typoClosings = data.Must(data.Each("typo-closings", 2, typoClosingOf))
