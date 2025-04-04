package main

import "fmt"

type character struct {
	name     string
	leftHand *item
}

type item struct {
	name string
}

func (c *character) pickup(i *item) {
	if i == nil {
		return
	}

	if c.leftHand != nil {
		fmt.Printf("%v picks down %v", c.name, c.leftHand.name)
	}

	fmt.Printf("%v picks up %v\n", c.name, i.name)
	c.leftHand = i
}

func (c *character) give(to *character) {
	if to == nil || to == nil {
		return
	}

	if c.leftHand == nil {
		fmt.Printf("%v has nothing to give\n", c.name)
		return
	}

	if to.leftHand != nil {
		fmt.Printf("%v hands are full\n", to.name)
		return
	}

	fmt.Printf("%v gives %v to %v\n", c.name, c.leftHand.name, to.name)
	to.leftHand = c.leftHand
	c.leftHand = nil
}

func (c *character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}
