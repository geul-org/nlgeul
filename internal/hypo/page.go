//ff:func feature=hypo type=helper control=sequence
//ff:what Page — the combinations of candidates of one page of a sorted list and how many come before them; pages are counted from 1, page 0 is the whole list, a page past the last one is empty
package hypo

// Page cuts the list, it does not sort or score: the rank of the i-th
// combination it gives is the number before them + i + 1. PerPage is how
// many combinations a page holds and is at least 1.
func Page(listed []Hypothesis, page, perPage int) (int, []Hypothesis) {
	if page < 1 {
		return 0, listed
	}
	if page-1 > (len(listed)-1)/perPage || len(listed) == 0 {
		return len(listed), nil
	}
	first := (page - 1) * perPage
	if perPage > len(listed)-first {
		return first, listed[first:]
	}
	return first, listed[first : first+perPage]
}
