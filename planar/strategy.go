package planar

import (
	"sync"
)

var algorithmMegrez Algorithm
var once sync.Once

type newAlgorithm func() Algorithm

// NormalStrategy returns normal algorithm.
func NormalStrategy() Algorithm {
	return GetStrategy(newMegrezAlgorithm)
}

// GetStrategy returns  algorithm by new Algorithm.
func GetStrategy(f newAlgorithm) Algorithm {
	return f()
}

func newMegrezAlgorithm() Algorithm {
	once.Do(func() {
		algorithmMegrez = &MegrezAlgorithm{}
	})
	return algorithmMegrez
}
