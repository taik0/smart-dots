package main

import "fmt"

type Vector struct {
	X float64
	Y float64
}

func (v1 *Vector) Add(v2 Vector) {
	v1.X += v2.X
	v1.Y += v2.Y
}

func (v Vector) String() string {
	return fmt.Sprintf("X: %f, Y: %f", v.X, v.Y)
}
