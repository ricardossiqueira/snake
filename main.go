package main

import (
	"fmt"
	"snake/cli"
	"snake/food"
	"snake/snake"
	"time"

	"github.com/mattn/go-tty"
)

type game struct {
	f     *food.Food
	s     *snake.Snake
	score int
}

func (g *game) updateScore() {
	g.score = len(g.s.Body)
}

func (g *game) printScore() {
	g.updateScore()
	cli.MoveCursor(1, 1)
	fmt.Printf("Score: %d", g.score)
}

// func debugSnake(s snake.Snake) {
// 	cli.MoveCursor(1, 1)
// 	fmt.Print(s.Body)
// }

func handleInput(s *snake.Snake) error {
	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			return err
		}
		switch char {
		case 'A':
			s.Dir = snake.UP
		case 'B':
			s.Dir = snake.DOWN
		case 'C':
			s.Dir = snake.LEFT
		case 'D':
			s.Dir = snake.RIGHT
		}
	}
}

func main() {
	c, err := cli.New()
	if err != nil {
		panic("cli.New: failed to initialize the cli handler")
	}

	f := food.New()
	s := snake.New()
	g := &game{f, s, len(s.Body)}

	go handleInput(s)

	for s.Alive {
		cli.ClearCli()
		g.printScore()

		if !g.f.Alive {
			f.SpawnRand(c)
		}

		g.f.Draw()

		g.s.HandleInput()
		g.s.Eat(f)

		g.s.Draw()
		g.s.HitWall(c)
		g.s.HitSelf()
		time.Sleep(time.Millisecond * 100)
	}

	cli.MoveCursor(0, 0)
	cli.ShowCursor()
	fmt.Printf("Game over! Score:%d\n", g.score)
}
