//ff:func feature=parse type=rule control=sequence
//ff:what newIndex — build the token index (an error for a spacing rule no name of SpacingArticles gives): the closed-set words, then a graph for every candidate of every particle, closing and slot of the tables, the graphs of the signs, the graph of a family name, the graphs of a new word and of a stem's restored syllable, and with a word table the particles read against the final consonant of the word before them (addMisfit), the markers written as they are spoken (addSpoken), its keys, the wrong ways to write a closed-set word, the closings written with the ㅎ of a stem's last 하 fused into them (addHa), and a graph for each kind of stem, and the keys longer than a new word's stem may be (addFarKeys); with a word table the rules of a sentence, and a graph the case exception table names that is neither a marker graph nor a rule of a sentence of the index is an error (checkCases)
package parse

// The graphs of a new word and of a stem's restored syllable are
// addStems'.
func newIndex(o Options) (*index, error) {
	spacing, err := spacingSet(o.Spacing)
	if err != nil {
		return nil, err
	}
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, lexed: o.Words != "", spacing: spacing}
	ix.addWords()
	ix.addParticles()
	ix.addClosings()
	ix.addPicked()
	ix.addChains()
	ix.addNamed()
	ix.addSlots()
	if o.Words != "" {
		ix.addWide()
		ix.addMisfit()
		ix.addSpoken()
	}
	ix.addCopula()
	ix.addFused()
	ix.addNouns()
	ix.addDots()
	ix.addPaths()
	ix.addPairs()
	ix.addSigns()
	ix.addNumbers()
	ix.addSurname()
	ix.addStems()
	if o.Words == "" {
		ix.addFarKeys()
		ix.addEnds()
		return ix, nil
	}
	lex, typos, total, err := loadWords(o.Words)
	if err != nil {
		return nil, err
	}
	ix.lex, ix.typos, ix.lexTotal, ix.layers = lex, typos, total, layerCost
	ix.addTypoKeys()
	ix.addAlts()
	ix.addHa()
	if o.Layers >= layerKnown && o.Layers <= layerCost {
		ix.layers = o.Layers
	}
	ix.addKinds()
	ix.addFarKeys()
	ix.addEnds()
	ix.addSentence()
	if err := ix.checkCases(); err != nil {
		return nil, err
	}
	return ix, nil
}
