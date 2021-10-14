package planar

import (
	"sync"
)

var algorithmMegrez Algorithm
var once sync.Once

type newAlgorithm func() Algorithm

// NormalStrategy returns normal algorithm.
func NormalStrategy() Algorithm {
	return GetStrategy(NewMegrezAlgorithm)
}

// GetStrategy returns  algorithm by new Algorithm.
func GetStrategy(f newAlgorithm) Algorithm {
	return f()
}

//  NewMegrezAlgorithm returns Algorithm that is MegrezAlgorithm.
func NewMegrezAlgorithm() Algorithm {
	once.Do(func() {
		algorithmMegrez = &megrezAlgorithm{}
	})
	return algorithmMegrez
}
