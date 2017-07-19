// $ go run time_channels.go
package main

import (
	"fmt"
	"os"
	"time"
)

var boolmap map[int]bool

func init() {
	boolmap = make(map[int]bool)

	for i := 0; i < 10; i++ {
		boolmap[i] = false
	}
	printMap(boolmap)
}

func main() {
	ticker := time.NewTicker(time.Second * 3)
	timer := time.NewTimer(time.Second * 60).C
	doner := make(chan bool)

	go func() {
		inc := 0
		for t := range ticker.C {
			boolmap[inc] = true
			inc++
			fmt.Println("Tick at", t)
			printMap(boolmap)
			if allTrue(boolmap) {
				doner <- true
				break
			}
		}
	}()

	for {
		select {
		case <-timer:
			fmt.Println("Timer expired.")
			ticker.Stop()
			os.Exit(1)
		case <-doner:
			fmt.Println("All done!")
			ticker.Stop()
			os.Exit(0)
		}
	}
}

func printMap(mapVar map[int]bool) {
	for i := range mapVar {
		fmt.Println(i, mapVar[i])
	}
}

func allTrue(mapVar map[int]bool) bool {
	mapStatus := false
	for i := range mapVar {
		if mapVar[i] == false {
			mapStatus = false
			break
		} else {
			mapStatus = true
		}
	}

	return mapStatus
}

// References:
// https://mmcgrana.github.io/2012/09/go-by-example-timers-and-tickers.html
// https://gobyexample.com/timers
// https://gobyexample.com/goroutines
