package main

import (
	"image"
	"strconv"
	"time"
)

func main() {

	//// ================ gps.go ================
	//mars := world{radius: 3389.5}
	//bradbury := location{
	//	latitude:  -4.5895,
	//	longitude: 137.4417,
	//	name:      " Bradbury Landing",
	//}
	//elysium := location{
	//	latitude:  4.5,
	//	longitude: 135.9,
	//	name:      "Elysium Planitia",
	//}
	//gps := gps{
	//	current:     bradbury,
	//	destination: elysium,
	//	world:       mars,
	//}
	//curiosity := rovers{
	//	gps,
	//}
	//
	//fmt.Println(curiosity.message())
	//
	//// ================ marshal.go ================
	//elysium2 := location2{
	//	Name: "Elysium Planitia",
	//	Lat:  coordinate{4, 30, 0.0, 'N'},
	//	Long: coordinate{135, 54, 0.0, 'E'},
	//}
	//
	//bytes, err := json.MarshalIndent(elysium2, "", " ")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(string(bytes))
	//
	//// ================ animals.go ================
	//var sol, hour int
	//const sunrise, sunset = 8, 18
	//animals := []animal{
	//	rabbit{name: "Roger Rabbit"},
	//	fish{name: "Roger Fish"},
	//}
	//
	//for {
	//	fmt.Sprintf("%2d:00 ", hour)
	//
	//	if hour < sunrise || hour > sunset {
	//		fmt.Println("Animals are sleeping")
	//	} else {
	//		i := rand.Intn(len(animals))
	//		step(animals[i])
	//	}
	//
	//	time.Sleep(500 * time.Millisecond)
	//	hour++
	//
	//	if hour >= 24 {
	//		hour = 0
	//		sol++
	//		if sol >= 3 {
	//			break
	//		}
	//	}
	//
	//}
	//
	//// ================ turtle.go ================
	//var turtle turtle
	//turtle.moveDown()
	//turtle.moveDown()
	//turtle.moveDown()
	//turtle.moveRight()
	//fmt.Println(turtle.String())
	//
	//// ================ knights.go ================
	//lancelot := &character{name: "Lancelot"}
	//king := &character{name: "King Arthur"}
	//sword := &item{name: "Sword"}
	//axe := &item{name: "Axe"}
	//lancelot.pickup(sword)
	//lancelot.give(king)
	//lancelot.give(king)
	//lancelot.pickup(axe)
	//lancelot.give(king)
	//
	//// ================ sudoku.go ================
	//sudoku := newSudoku([9][9]int8{
	//	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	//	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	//	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	//	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	//	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	//	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	//	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	//	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	//	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	//})
	//err = sudoku.place(7, 6, 7)
	//
	//if err != nil {
	//	var errs SudokuError
	//	if errors.As(err, &errs) {
	//		fmt.Printf("Sudoku placement failed with %d errors\n", len(errs))
	//		for _, e := range errs {
	//			fmt.Printf("- %v\n", e)
	//		}
	//	}
	//}
	//
	//err = sudoku.place(2, 0, 9)
	//
	//// ================ select.go ================
	//c := make(chan int)
	//for i := 0; i < 5; i++ {
	//	go sleepyGopher(i, c)
	//}
	//timeout := time.After(2 * time.Second)
	//for i := 0; i < 5; i++ {
	//	select {
	//	case gopherId := <-c:
	//		fmt.Println("gopherId ", gopherId)
	//	case <-timeout:
	//		fmt.Println("timeout")
	//	}
	//}
	//
	//// ================ pipeline.go ================
	//c0 := make(chan string)
	//c1 := make(chan string)
	//go sourceGopher(c0)
	//go filterGopher(c0, c1)
	//printGopher(c1)
	//
	//// ================ pipeline.go ================
	//c3 := make(chan string)
	//c4 := make(chan string)
	//go sourceDuplicatesGopher(c3)
	//go removeDuplicates(c3, c4)
	//printerDuplicatesGopher(c4)
	//
	//// ================ scrape.go ================
	//urls := []string{
	//	"https://www.google.com",
	//	"https://www.openai.com",
	//	"https://www.stackoverflow.com",
	//	"https://www.github.com",
	//	"https://www.example.com",
	//}
	//var urlsToVisit []string
	//for i := 0; i < 10000; i++ {
	//	urlsToVisit = append(urlsToVisit, urls[i%len(urls)])
	//}
	//var wg sync.WaitGroup
	//visited := &Visited{visited: make(map[string]int)}
	//for _, url := range urlsToVisit {
	//	wg.Add(1)
	//	go visited.VisitLink(url, &wg)
	//}
	//wg.Wait()
	//
	//for url, count := range visited.visited {
	//	fmt.Printf("visited: %v, count: %v\n", url, count)
	//}

	// ================ rover.go ================
	grid := NewMarsGrid(image.Point{X: 100, Y: 100})
	rovers := make([]*RoverDriver, 5)
	for i := range rovers {
		rovers[i] = deployRover("Rover "+strconv.Itoa(i), grid)
		rovers[i].Start()
	}
	time.Sleep(60 * time.Second)
}
