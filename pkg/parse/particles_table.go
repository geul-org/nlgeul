package parse

import (
	"github.com/geul-org/nlgeul/data"
)

// particles are the particles of the table particles.tsv in the data
// package (from particles.md sections 1 to 4), with the counts written
// there (1 where the table says observed without a number). Left out: forms with no observation, forms that are also a
// stem plus an ending (하고, 보고, 치고, 말고, 다가, 마는), the
// quoting 고 and 이다 (read as endings), and the one-consonant forms.
// The third field is the particle's place in a chain (section 5): 0 for
// the suffix-like 들, 씩, 쯤, 끼리, then 1 to 5. The fourth is the case
// the tables give the particle (0 for an auxiliary, a connective or a
// suffix); 에서 and 서 mark both the subject of a group and a place.
var particles = data.Must(data.Each(particlesFile, 5, particleOf))
