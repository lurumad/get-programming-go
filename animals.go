package main

import (
	"fmt"
	"math/rand"
)

type animal interface {
	move() string
	eat() string
}

type rabbit struct {
	name string
}

func (r rabbit) move() string {
	switch rand.Intn(2) {
	case 0:
		return "running"
	default:
		return "stepping"
	}
}

func (r rabbit) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "carrot"
	default:
		return "grass"
	}
}

type fish struct {
	name string
}

func (f fish) move() string {
	switch rand.Intn(2) {
	case 0:
		return "swims forward"
	default:
		return "swims backward"
	}
}

func (f fish) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "algae"
	default:
		return "plankton"
	}
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v. \n", a, a.move())
	default:
		fmt.Printf("%v %v. \n", a, a.eat())
	}
}
