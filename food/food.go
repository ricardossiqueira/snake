package food

import (
	"fmt"
	"math/rand"
	"snake/cli"
)

type Food struct {
	X, Y  int
	Alive bool
}

func New() *Food {
	return &Food{1, 1, false}
}

func (f *Food) SpawnRand(c *cli.Cli) {
	f.X = randPos(c.W)
	f.Y = randPos(c.H)
	f.Alive = true
}

func (f *Food) Draw() {
	cli.MoveCursor(f.X, f.Y)
	fmt.Print("*")
	// fmt.Print(f.X, f.Y)
}

func randPos(boundary int) int {
	return (rand.Intn(boundary) + 1)
}
