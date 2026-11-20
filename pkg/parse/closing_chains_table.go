package parse

import "github.com/geul-org/nlgeul/pkg/data"

// A sentence closed by a final closing and quoted with no quotation
// marks may take the ending of the quoting verb right after that closing,
// the quoting 고 and the verb 하- that stood between them dropped: 있다거나
// (있다고 하거나), 가리라는 (가리라고 하는), 학생이라기보다 (학생이라고
// 하기보다), 갑시다고; or the quoting 라고 하- shortened onto it: 간다라는,
// 간다라면, 가느냐라는. closingChains gives, by the name of such an
// ending (closing-chains.tsv, the first column), the names of the final
// closings it is read right after (the second, joined with |). The ending
// is read as one more piece of the tail of the final closing before it,
// so the stem is judged by that closing. A closed table: only the pairs
// listed are read, also where the two closings are written as one closing
// of the tables (다거나: as the gold of the evaluation mostly reads it).
// Each ending has a graph of its own (chainAuxName and the name), whose
// exception is that the closing before it is not one of those listed.
const chainAuxName = "사슬:맺음 뒤 맺음 "

var closingChains = data.Must(data.Map("closing-chains"))
