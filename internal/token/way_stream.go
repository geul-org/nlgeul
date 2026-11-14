//ff:type feature=token type=model
//ff:what wayStream — what the box output asks of a stream: the best way of the next stretch, the way after the one sent last for a stretch, and the trace lines of the units read (a Stream is one)
package token

type wayStream interface {
	Next() (Segment, bool)
	Again(Span) (Segment, bool)
	Notes() []string
}
