/*
  We have been asked to create a tool to help users manage their calendars. 
  Given an unordered list of times of day when someone is busy, 
  write a function that tells us whether they're available during a specified period of time.
  Each time is expressed as an integer using 24-hour notation, 
  such as 1200 (12:00), 1530 (15:30), or 800 (8:00).
 */

package main

import "fmt"

func main() {
  appointments := [][]int{
    {1230, 1300},
    {845, 900},
    {1300, 1500},
  }

    fmt.Println(isAvailable(appointments,  830,  845)  == true)
    fmt.Println(isAvailable(appointments, 1330, 1400)  == false)    
	fmt.Println(isAvailable(appointments,  830,  930)  == false)    
	fmt.Println(isAvailable(appointments,  855,  930)  == false)    
	fmt.Println(isAvailable(appointments, 1500, 1600)  == true)    
	fmt.Println(isAvailable(appointments,  845,  900)  == false)    
	fmt.Println(isAvailable(appointments, 1229, 1231)  == false)    
}

func isAvailable(appts [][]int, stime int, etime int) bool {
    for _, appt := range appts {
        astart := appt[0]
        aend := appt[1]
        if stime < astart && etime <= astart {
            continue
        } else if stime >= aend && etime > aend {
            continue
        } else {
            return false
        }
    }
    
    return true
}	
