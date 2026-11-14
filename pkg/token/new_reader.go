//ff:func feature=token type=reader control=sequence
//ff:what New — a reader with its index built, and its word table loaded when one is named
package token

func New(o Options) (*Reader, error) {
	ix, err := newIndex(o)
	if err != nil {
		return nil, err
	}
	return &Reader{ix: ix}, nil
}
