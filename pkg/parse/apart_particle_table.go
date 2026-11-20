package parse

import "slices"

// apartEnds are the kinds of the last piece of a word that a particle
// written apart after it cannot belong to (apartParticleRule): a
// particle, the copula, and a closing. apartBefore are the kinds of the
// name of a word right before such a particle that it may belong to
// although they are no nominal: a closing sign of a pair, a run of Latin
// letters and digits, a name joined with a hyphen and a code (a dash is
// told by its last piece: apartDash).
var (
	apartEnds   = slices.Concat(particleKinds, []string{"서술격"}, predKinds)
	apartBefore = []string{shutName, labelOf(runName), "붙임", codeName}
)
