package main

import (
	"fmt"
	"math"
)

type rovers struct {
	gps
}

type gps struct {
	current     location
	destination location
	world       world
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.1f km to %v", g.distance(), g.destination.description())
}

type location struct {
	latitude, longitude float64
	name                string
}

func (l location) description() string {
	return fmt.Sprintf("%+v\n", l)
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(radians(p1.latitude))
	s2, c2 := math.Sincos(radians(p2.latitude))
	clongitude := math.Cos(radians(p1.longitude - p2.longitude))
	return w.radius * math.Acos(s1*s2+c1*c2*clongitude)
}

func radians(deg float64) float64 {
	return deg * math.Pi / 180
}
