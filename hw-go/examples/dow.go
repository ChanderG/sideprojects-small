/* print day of week in a round about manner
  mainly as an excuse to try out switch
*/
package main

import "fmt"
import "time"

func main() {
  //fmt.Println("Time: ",time.Now().Weekday())
  switch time.Now().Weekday() {
	case  time.Saturday, time.Sunday:
	  fmt.Println("Weekend!!")
	case  time.Tuesday, time.Thursday, time.Friday:
	  fmt.Println("Exercise!!")
	default:
	  fmt.Println("Whatever")
  }

  if t := time.Now().Hour(); t < 12 {
	fmt.Println("AM")
  } else if t == 12 {
	fmt.Println("Noon")
  } else if t == 24 {
	fmt.Println("Midnight")
  } else {
	fmt.Println("PM")
  }
}
