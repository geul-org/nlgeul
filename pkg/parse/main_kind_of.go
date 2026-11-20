//ff:func feature=parse type=helper control=sequence
//ff:what mainKindOf — the name a stem an auxiliary verb shares is read under as a main predicate: the word table's kind written for it in mainKinds (있: the adjective), the verb for any other
package parse

func mainKindOf(stem string) string {
	if k, ok := mainKinds[stem]; ok {
		return keyKind + ":" + kindLabel[k]
	}
	return keyKind + ":" + kindLabel[kindVerb]
}
