package token

// nameJoins are the signs that join the parts of a name a path is
// written with (a-b, a_b). Hangul syllables right after one of them at a
// path's end are a part of the name (files/a-가격정책), as a name part
// right after a slash or a dot is; Hangul syllables written right onto a
// Latin letter or a digit are a particle or a word of the sentence (a/b를).
const nameJoins = "-_"
