package main

import (
	"fmt"
)

func sourceDuplicatesGopher(downstream chan string) {
	for _, text := range []string{"a", "b", "b", "c", "d", "d", "d", "e"} {
		downstream <- text
	}
	close(downstream)
}

func removeDuplicates(upstream, downstream chan string) {
	prev := ""
	for item := range upstream {
		if prev != item {
			prev = item
			downstream <- item
		}
	}
	close(downstream)
}

func printerDuplicatesGopher(upstream chan string) {
	for item := range upstream {
		fmt.Println(item)
	}
}
