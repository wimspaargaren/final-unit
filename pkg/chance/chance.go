// Package chance provides functionality for easily dealing with chances
package chance

import (
	"math/rand"
)

const fullChance = 100

// IsChance returns booleans for given chance
func IsChance(chance float64) bool {
	if chance >= 0 && chance <= fullChance {
		temp := chance / fullChance
		return rand.Float64() < temp
	}
	return false
}

// GetIndex get index based on given length
func GetIndex(length int) int {
	if length == 0 {
		return 0
	}
	return rand.Intn(length)
}
