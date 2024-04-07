package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

type Cli struct {
	W, H int
}

const (
	clearCli   = "\033[2J"
	hideCursor = "\033[?25l"
	showCursor = "\033[?25h"
	moveCursor = "\033[%d;%dH"
)

func New() (*Cli, error) {
	h, w, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	// handle ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ShowCursor()
		MoveCursor(0, 0)
		os.Exit(1)
	}()

	// setup term screen
	HideCursor()
	ClearCli()

	return &Cli{W: w, H: h}, nil
}

func ClearCli() {
	fmt.Print(clearCli)
}

func HideCursor() {
	fmt.Print(hideCursor)
}

func ShowCursor() {
	fmt.Print(showCursor)
}

func MoveCursor(x, y int) {
	fmt.Printf(moveCursor, x, y)
}
