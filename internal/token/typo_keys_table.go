package token

// typoKeys is the closed table of the wrong ways to write a word of
// a closed set that this layer reads (a written form to the word it
// stands for): 꺼 for the dependent noun 거. The table is closed and
// this line is all of it; it is not moved from any other list or
// dictionary. It is read when there is a word table (addTypoKeys).
var typoKeys = map[string]string{"꺼": "거"}
