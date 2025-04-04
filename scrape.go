package main

import "sync"

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
}
