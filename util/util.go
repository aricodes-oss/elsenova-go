package util

import (
	"fmt"
	"math/rand/v2"
)

// Any positive 10 digit integer
const SEED_MAX = 9999999999

func Map[I any, R any](input []I, mapper func(I, int) R) []R {
	results := make([]R, len(input))

	for idx, elem := range input {
		results[idx] = mapper(elem, idx)
	}

	return results
}

func RandoSeed() string {
	return fmt.Sprintf("%010d", rand.IntN(SEED_MAX))
}
