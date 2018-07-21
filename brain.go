package main

import (
	"math"
	"math/rand"
)

type Brain struct {
	directions []Vector
	step       int
}

func (b *Brain) Clone() *Brain {
	newDirections := make([]Vector, len(b.directions))
	for i, v := range b.directions {
		newDirections[i] = v
	}
	return &Brain{directions: newDirections}
}

func (b *Brain) Mutate() {
	var mutationRate float64 = 0.01
	for i := range b.directions {
		if rand.Float64() < mutationRate {
			randomAngle := rand.Float64() * (2 * math.Pi)
			b.directions[i] = vectorFromAngle(randomAngle)
		}
	}
}

func NewBrain(size int) *Brain {
	directions := make([]Vector, size)
	for i := range directions {
		randomAngle := rand.Float64() * (2 * math.Pi)
		directions[i] = vectorFromAngle(randomAngle)
	}
	return &Brain{directions: directions}
}
