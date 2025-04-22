package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
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
	name     string
	commandc chan command
	occupier *Occupier
}

func (rover *RoverDriver) drive() {
	fmt.Printf("%s deployed at initial position %v", rover.name, rover.occupier.Position())
	stopped := true
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-rover.commandc:
			switch c {
			case right:
				direction = image.Point{X: -direction.Y, Y: direction.X}
				log.Printf("%s turns to the right, new direction: %v\n", rover.name, direction)
			case left:
				direction = image.Point{X: direction.Y, Y: -direction.X}
				log.Printf("%s turns to the left, new direction: %v\n", rover.name, direction)
			case start:
				stopped = false
				log.Printf("%s starting", rover.name)
			case stop:
				stopped = true
				log.Printf("%s stopping", rover.name)
			}

		case <-nextMove:
			if !stopped {
				nextMove = time.After(updateInterval)
				newPositon := rover.occupier.Position().Add(direction)
				if rover.occupier.Move(newPositon) {
					log.Printf("%s moved to %v", rover.name, newPositon)
					break
				}
				log.Printf("%s blocked trying to move from %v to %v", rover.name, rover.occupier.Position(), newPositon)
			}
		}
	}
}

func (rover *RoverDriver) randomPositon() {

}

func (rover *RoverDriver) Left() {
	rover.commandc <- left
}

func (rover *RoverDriver) Right() {
	rover.commandc <- right
}

func (rover *RoverDriver) Start() {
	rover.commandc <- start
}

func (rover *RoverDriver) Stop() {
	rover.commandc <- stop
}

func NewRoverDriver(name string, occupier *Occupier) *RoverDriver {
	r := &RoverDriver{
		name:     name,
		commandc: make(chan command),
		occupier: occupier,
	}
	go r.drive()
	return r
}

type cell struct {
	occupier *Occupier
}

type MarsGrid struct {
	bounds image.Rectangle
	mutex  sync.Mutex
	cells  [][]cell
}

func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{Max: size},
		cells:  make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
	}
	return grid
}

func (grid *MarsGrid) Size() image.Point {
	return grid.bounds.Max
}

func (grid *MarsGrid) Occupy(point image.Point) *Occupier {
	grid.mutex.Lock()
	defer grid.mutex.Unlock()
	cell := grid.cell(point)

	if cell == nil || cell.occupier != nil {
		return nil
	}

	return &Occupier{grid: grid, position: point}
}

func (grid *MarsGrid) cell(point image.Point) *cell {
	if !point.In(grid.bounds) {
		return nil
	}

	return &grid.cells[point.Y][point.X]
}

type Occupier struct {
	grid     *MarsGrid
	position image.Point
}

func (occupier *Occupier) Move(point image.Point) bool {
	occupier.grid.mutex.Lock()
	defer occupier.grid.mutex.Unlock()
	cell := occupier.grid.cell(point)
	if cell == nil || cell.occupier != nil {
		return false
	}
	occupier.grid.cell(occupier.position).occupier = nil
	cell.occupier = occupier
	occupier.position = point
	return true
}

func (occupier *Occupier) Position() image.Point {
	return occupier.position
}

func deployRover(name string, grid *MarsGrid) *RoverDriver {
	var occupier *Occupier
	for occupier == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		occupier = grid.Occupy(startPoint)
	}
	return NewRoverDriver(name, occupier)
}
