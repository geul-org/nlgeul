package token

// spoken are the closed-set words that are another word as it is said
// in speech: 거 is the dependent noun 것 (거야, 할 거다). A word table
// counted from written text has few of the spoken form and many of the
// word itself, so that the form was rarely seen there does not tell
// against it. A closed table: only what is listed is read so.
var spoken = wordSets[setSpoken]
