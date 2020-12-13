package geos

const GEOS string = "GEOS"

func NormalStrategy() Algorithm {
	return AlgorithmStrategy(GEOS)
}

func AlgorithmStrategy(name string) Algorithm {
	switch name {
	case GEOS:
		return new(GEOSAlgorithm)
	default:
		return nil
	}
}
