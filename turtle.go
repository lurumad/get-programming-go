package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) moveUp() {
	t.y--
}

func (t *turtle) moveDown() {
	t.y++
}

func (t *turtle) moveLeft() {
	t.x--
}

func (t *turtle) moveRight() {
	t.x++
}

func (t *turtle) String() string {
	return fmt.Sprintf("%v %v", t.x, t.y)
}
