package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, text := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- text
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}
