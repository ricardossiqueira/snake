package snake

import (
	"fmt"
	"snake/cli"
	"snake/food"
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

type BodyElem struct{ X, Y int }

type Snake struct {
	Alive bool
	Body  []BodyElem
	Dir   string
}

func New() *Snake {
	s := &Snake{Alive: true, Body: []BodyElem{{10, 10}}}

	return s
}

func (s *Snake) Draw() {
	for i, pos := range s.Body {
		cli.MoveCursor(pos.X, pos.Y)
		if i == 0 {
			fmt.Print("O")
		} else {
			fmt.Print("o")
		}
		// fmt.Print(pos.X, pos.Y)
	}
}

func (s *Snake) HandleInput() {
	switch s.Dir {
	case UP:
		s.MoveUp()
	case DOWN:
		s.MoveDown()
	case LEFT:
		s.MoveLeft()
	case RIGHT:
		s.MoveRight()
	}

}

func (s *Snake) Eat(f *food.Food) {
	l := len(s.Body)
	t := s.Body[l-1]
	h := s.Body[0]
	if h.X == f.X && h.Y == f.Y {
		f.Alive = false
		s.Body = append(s.Body, BodyElem{t.X, t.Y})
	}
}

func (s *Snake) HitWall(c *cli.Cli) {
	for i, pos := range s.Body {
		if pos.Y < 0 {
			s.Body[i].Y = c.H
		}
		if pos.Y > c.H {
			s.Body[i].Y = 0
		}
		if pos.X < 0 {
			s.Body[i].X = c.W
		}
		if pos.X > c.W {
			s.Body[i].X = 0
		}
	}
}

func (s *Snake) HitSelf() {
	h := s.Body[0]
	if len(s.Body) > 2 {
		for _, pos := range s.Body[1:] {
			if h == pos {
				s.Alive = false
			}
		}
	}
}

func (s *Snake) MoveRight() {
	s.MoveTail()
	s.Body[0].Y--
}

func (s *Snake) MoveLeft() {
	s.MoveTail()
	s.Body[0].Y++
}

func (s *Snake) MoveUp() {
	s.MoveTail()
	s.Body[0].X--
}

func (s *Snake) MoveDown() {
	s.MoveTail()
	s.Body[0].X++
}

func (s *Snake) MoveTail() {
	l := len(s.Body)
	s.Body = append(s.Body[:1], s.Body[:l]...)[:l]
}
