package parse

import (
	"slices"
)

// argumentPronouns are the pronouns of the closed sets whose letters with
// a particle stand as an argument of their own (pronounParticled):
// the demonstratives 이, 그, 저 and the personal pronouns. particleForms
// are the forms of the particles of the particle tables.
var (
	argumentPronouns = slices.Concat(pronouns, personalPronouns)
	particleForms    = markerForms(ParticleMarkers())
)
