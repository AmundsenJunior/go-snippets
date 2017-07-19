package main

import (
	"fmt"
	"os"
	//"sync"
	"time"
)

var boolmap map[int]bool

//var wg sync.WaitGroup
var inc int

func init() {
	inc = 0

	boolmap = make(map[int]bool)
	for i := 0; i < 10; i++ {
		boolmap[i] = false
	}
	//printMap(boolmap)
}

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	ticker := time.NewTicker(time.Second * 3).C
	timer := time.NewTimer(time.Second * 40).C
	done := make(chan bool)

	defer fmt.Println("Deferred statement.")

	for {
		select {
		case <-ticker:
			//wg.Add(1)
			if allTrue(boolmap) {
				fmt.Println("all true")
				done <- true
				return
			}
			fmt.Printf("Ticker at %v.\n", inc)
			go routine(boolmap, inc)
			inc++
		case <-timer:
			fmt.Println("Timer expired.")
			//ticker.Stop()
			//wg.Wait()
			exitCode = 1
			break
			// case <-done:
			// 	fmt.Println("All done!")
			// 	//ticker.Stop()
			// 	//wg.Wait()
			// 	exitCode = 0
			// 	break
		}
	}
	fmt.Println("All done!")
	exitCode = 0
	<-done
}

func routine(boolmap map[int]bool, inc int) {
	//defer wg.Done()

	boolmap[inc] = true
	printMap(boolmap)
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
