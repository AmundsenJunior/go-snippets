// $ go run ticker.go
package main

import (
	"fmt"
	"time"
)

var boolmap map[int]bool

func init() {
	boolmap = make(map[int]bool)

	for i := 0; i < 10; i++ {
		boolmap[i] = false
	}
}

func main() {
	ticker := time.NewTicker(time.Second * 5)

	go func() {
		inc := 0
		for t := range ticker.C {
			boolmap[inc] = true
			inc++
			fmt.Println("Tick at", t)
			printMap(boolmap)
		}
	}()

	time.Sleep(time.Second * 61)
	ticker.Stop()
	fmt.Println("Ticker stopped.")
}

func printMap(mapVar map[int]bool) {
	for i := range mapVar {
		fmt.Println(i, mapVar[i])
	}
}

// References:
// https://gobyexample.com/tickers
// https://gobyexample.com/maps
