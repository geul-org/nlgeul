package token

import "github.com/geul-org/nlgeul/pkg/rule"

// typoClosings are the closings that are often written wrong, each with
// the right closing it stands for. The rows are common misspellings
// written down from memory, not copied from another list or a
// dictionary, and this is the whole table: a closing written wrong in
// some other way is not read. A row whose Form has a Lead is the written
// form under the final consonant ㄹ (할께 for 할게: 하 + ㄹ께); a row of
// whole syllables is read after the stems the right closing is read
// after (먹을께 for 먹을게, 했읍니다 for 했습니다). The right closing is
// found by its name among the closings of the tables; every row must
// find one. Not in the table, because the right form is not a closing of
// the tables: ㄹ려는 for 려는, ㄹ래야 for 려야, ㄹ런지 and ㄹ른지 for
// ㄹ는지, 는구만 for 는구먼.
var typoClosings = []typoClosing{
	{"께", rule.JongL, "맺음:ㄹ게"},
	{"을께", 0, "맺음:을게"},
	{"껄", rule.JongL, "맺음:ㄹ걸"},
	{"을껄", 0, "맺음:을걸"},
	{"찌", rule.JongL, "맺음:ㄹ지"},
	{"을찌", 0, "맺음:을지"},
	{"쑤록", rule.JongL, "맺음:ㄹ수록"},
	{"을쑤록", 0, "맺음:을수록"},
	{"려고", rule.JongL, "맺음:려고"},
	{"을려고", 0, "맺음:으려고"},
	{"려면", rule.JongL, "맺음:려면"},
	{"을려면", 0, "맺음:으려면"},
	{"읍니다", 0, "맺음:습니다"},
	{"슴", 0, "이름:음"},
	{"구만", 0, "맺음:구먼"},
	{"예요", 0, "맺음:에요"},
	{"에요", 0, "맺음:예요"},
}
