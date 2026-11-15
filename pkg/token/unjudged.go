//ff:func feature=token type=helper control=iteration dimension=1
//ff:what unjudged — the regions at the end of the queue that the rules of a sentence have not judged yet (the regions are judged in the order they are queued, so they are the last ones)
package token

func unjudged(queue []*region) []*region {
	from := len(queue)
	for from > 0 && !queue[from-1].judged {
		from--
	}
	return queue[from:]
}
