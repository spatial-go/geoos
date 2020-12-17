package geos

<<<<<<< HEAD
// GEOS geos algorithm name.
const GEOS string = "GEOS"

// NormalStrategy returns normal algorithm.
=======
// GEOS ...
const GEOS string = "GEOS"

// NormalStrategy ...
>>>>>>> a1310ea0c8d9c74dab601390b4f91981dc57f2c9
func NormalStrategy() Algorithm {
	return AlgorithmStrategy(GEOS)
}

<<<<<<< HEAD
// AlgorithmStrategy returns algorithm by name
=======
// AlgorithmStrategy ...
>>>>>>> a1310ea0c8d9c74dab601390b4f91981dc57f2c9
func AlgorithmStrategy(name string) Algorithm {
	switch name {
	case GEOS:
		return new(GEOSAlgorithm)
	default:
		return new(GEOSAlgorithm)
	}
}
