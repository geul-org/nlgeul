package token

// adnOnto are the dependent nouns whose stem the one-syllable rebuttal
// is lifted for when they are written right onto the modifying closing
// they stand on (closed-words-more.tsv): 것 and its spoken form 거
// (하는거, 없는게, 할것이). The other dependent nouns of the closed set
// are not in it: their letters are also a noun or a piece of a closing
// (수 the noun of number written onto a noun, 데 of the closings 는데 and
// ㄴ데), and the word before them may be read as ending in a modifying
// closing by a candidate other than the first (adnEnds asks any). Measured
// on the answer sheet they gave more wrong words than right: with every
// dependent noun of the closed set 8 more wrong and 3 more right than with
// these two, and with those the grammar tables read written onto a
// modifying closing (수 줄 데 때 채 too) 5 more wrong and none more right.
// A closed table: only what is listed is read so.
var adnOnto = wordSets[setAdnBound]
