package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

var (
	dotImg *ebiten.Image
	obsImg *ebiten.Image
	goal   Vector
)

func init() {
	dotImg, _ = ebiten.NewImage(2, 2, ebiten.FilterNearest)
	obsImg, _ = ebiten.NewImage(200, 5, ebiten.FilterNearest)
	goal = Vector{X: 200, Y: 20}
}

type Goal Vector

func (g Goal) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	dotImg.Fill(color.RGBA{0xff, 0xf5, 0x05, 0xff})
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)
	op.GeoM.Translate(g.X, g.Y)
	screen.DrawImage(dotImg, op)
}

type Dot struct {
	brain   *Brain
	pos     Vector
	vel     Vector
	acc     Vector
	dead    bool
	goal    bool
	fitness float64
	best    bool
}

func (d *Dot) Pos() (float64, float64) {
	return d.pos.X, d.pos.Y
}

func (d *Dot) String() string {
	return fmt.Sprintf("pos: %s, vel: %s, acc: %s ", d.pos, d.vel, d.acc)
}

func (d *Dot) SetPos(v Vector) {
	d.pos = v
}

func (d *Dot) SetAcc(v Vector) {
	d.acc = v
}

func (d *Dot) Update() {
	if !d.dead && !d.goal {
		d.move()
		if d.checkOOB(screenWidth, screenHeight) {
			d.dead = true
		}
		if d.goalDistance() < 5 {
			d.goal = true
		}
	}
}

func (d *Dot) checkOOB(width, height float64) bool {
	if d.pos.X < 4 || d.pos.Y < 4 || d.pos.X > width-4 || d.pos.Y > height-4 {
		return true
	}
	// if d.pos.X < 300 && d.pos.Y < 205 && d.pos.X > 100 && d.pos.Y > 200 {
	// 	return true
	// }
	return false
}

func (d *Dot) move() {
	if len(d.brain.directions) < d.brain.step {
		d.dead = true
		return
	}

	d.acc = d.brain.directions[d.brain.step]
	d.brain.step++
	d.vel.Add(d.acc)
	if d.vel.X > 2 {
		d.vel.X = 2
	}
	if d.vel.Y > 2 {
		d.vel.Y = 2
	}
	d.pos.Add(d.vel)
}

func (d Dot) goalDistance() float64 {
	return distance(d.pos, goal)
}

func (d *Dot) caltulateFitness() {
	if d.goal {
		d.fitness = (1.0/16.0 + 10000) / float64(d.brain.step*d.brain.step)
		return
	}
	d.fitness = 1.0 / (d.goalDistance() * d.goalDistance())
}

func (d Dot) Clone() *Dot {
	newDot := NewDot(Vector{X: screenWidth / 2, Y: screenHeight - 10})
	newDot.brain = d.brain.Clone()
	return newDot
}

func (d Dot) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	if d.dead {
		dotImg.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	} else {
		dotImg.Fill(color.White)
	}
	if d.best {
		dotImg.Fill(color.RGBA{0xff, 0x00, 0xff, 0xff})
	}
	op.ColorM.Scale(200.0/255.0, 200.0/255.0, 200.0/255.0, 1)
	op.GeoM.Translate(d.Pos())
	screen.DrawImage(dotImg, op)
}

func NewDot(pos Vector) *Dot {
	return &Dot{pos: pos, brain: NewBrain(1000)}
}

func vectorFromAngle(angle float64) Vector {
	return Vector{X: math.Cos(angle), Y: math.Sin(angle)}
}

func distance(v1 Vector, v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2))
}
