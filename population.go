package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type Population struct {
	dots       []*Dot
	fitnessSum float64
	generation int
	bestDot    int
	minStep    int
}

func (p Population) Draw(screen *ebiten.Image) {
	for i := range p.dots {
		p.dots[i].Draw(screen)
	}
}

func (p *Population) Update() {
	for i := range p.dots {
		if p.dots[i].brain.step > p.minStep {
			p.dots[i].dead = true
		} else {
			p.dots[i].Update()
		}
	}
}

func (p *Population) calculateFitness() {
	for i := range p.dots {
		p.dots[i].caltulateFitness()
	}
}

func (p Population) allDotsDead() bool {
	for i := range p.dots {
		if !p.dots[i].dead && !p.dots[i].goal {
			return false
		}
	}
	return true
}

func (p *Population) calculateFitnessSum() {
	p.fitnessSum = 0
	for i := range p.dots {
		p.fitnessSum += p.dots[i].fitness
	}
}

func (p *Population) setBestBot() {
	var max float64
	var maxIndex int
	for i := range p.dots {
		if p.dots[i].fitness > max {
			max = p.dots[i].fitness
			maxIndex = i
		}
	}
	p.bestDot = maxIndex

	if p.dots[maxIndex].goal {
		p.minStep = p.dots[maxIndex].brain.step
	}
}

func (p *Population) naturalSelection() {
	newDots := make([]*Dot, len(p.dots))
	p.setBestBot()
	p.calculateFitnessSum()
	newDots[0] = p.dots[p.bestDot].Clone()
	newDots[0].best = true
	for i := 1; i < len(p.dots); i++ {
		parent := p.selectParent()
		newDots[i] = parent.Clone()
	}
	p.dots = newDots
	p.generation += 1
}

func (p Population) selectParent() *Dot {
	n := rand.Float64() * p.fitnessSum
	var runningSum float64
	for i := range p.dots {
		runningSum += p.dots[i].fitness
		if runningSum > n {
			return p.dots[i]
		}
	}
	return nil
}

func (p *Population) mutate() {
	for i := range p.dots {
		p.dots[i].brain.Mutate()
	}
}

func NewRandomPopulation(size int) *Population {
	dots := make([]*Dot, size)
	for i := range dots {
		dots[i] = NewDot(Vector{X: screenWidth / 2, Y: screenHeight - 10})
	}
	return &Population{dots: dots, minStep: 1000, generation: 1}
}
