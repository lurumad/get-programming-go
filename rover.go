package main

import (
	"image"
	"log"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	stop  = command(2)
	start = command(3)
)

type RoverDriver struct {
	commandc chan command
}

func (r *RoverDriver) drive() {
	stopped := true
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
				log.Printf("new direction: %v\n", direction)
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
				log.Printf("new direction: %v\n", direction)
			case start:
				stopped = false
				log.Printf("starting")
			case stop:
				stopped = true
				log.Printf("stopping")
			}

		case <-nextMove:
			if !stopped {
				pos = pos.Add(direction)
				log.Printf("next move: %v\n", pos)
			}
			nextMove = time.After(updateInterval)
		}
	}
}

func (r *RoverDriver) Left() {
	r.commandc <- left
}

func (r *RoverDriver) Right() {
	r.commandc <- right
}

func (r *RoverDriver) Start() {
	r.commandc <- start
}

func (r *RoverDriver) Stop() {
	r.commandc <- stop
}

func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}
