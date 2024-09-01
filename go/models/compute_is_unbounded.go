package models

func computeIsBounded(particle Particle) (res bool) {

	switch v := particle.(type) {
	case OccurrenceDefinition:
		if v.GetIsLocalyUnbounded() {
			return true
		} else {
			// check if outer particle is unbounded

		}
	}

	return
}
