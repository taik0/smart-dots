package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

var (
	dots    *Population
	goalDot Goal
)

func update(screen *ebiten.Image) error {
	screen.Fill(color.Black)
	if dots.allDotsDead() {
		dots.calculateFitness()
		dots.naturalSelection()
		dots.mutate()
	} else {
		dots.Update()
		dots.Draw(screen)
		goalDot.Draw(screen)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Generation: %d\n", dots.generation))
	if dots.generation > 0 {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("\nBest: %d, minSteps: %d\n", dots.bestDot, dots.minStep))
	}
	return nil
}

func main() {
	goalDot = Goal{X: goal.X, Y: goal.Y}
	//obstacle = Obstacle{X: 100, Y: 200}
	rand.Seed(time.Now().Unix())
	dots = NewRandomPopulation(250)
	err := ebiten.Run(update, screenWidth, screenHeight, 2, "Ebiten")
	if err != nil {
		panic(err)
	}
}
