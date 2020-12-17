package geos

// GEOS ...
const GEOS string = "GEOS"

// NormalStrategy ...
func NormalStrategy() Algorithm {
	return AlgorithmStrategy(GEOS)
}

// AlgorithmStrategy ...
func AlgorithmStrategy(name string) Algorithm {
	switch name {
	case GEOS:
		return new(GEOSAlgorithm)
	default:
		return nil
	}
}
