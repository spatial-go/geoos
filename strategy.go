package geos

// GEOS geos algorithm name.
const GEOS string = "GEOS"

// NormalStrategy returns normal algorithm.
func NormalStrategy() Algorithm {
	return AlgorithmStrategy(GEOS)
}

// AlgorithmStrategy returns algorithm by name
func AlgorithmStrategy(name string) Algorithm {
	switch name {
	case GEOS:
		return new(GEOSAlgorithm)
	default:
		return new(GEOSAlgorithm)
	}
}
