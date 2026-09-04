package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	targetDate := time.Date(current.Year()+1, time.January, 1, 0, 0, 0, 0, current.Location())
	
	fmt.Printf("%.2f\n", targetDate.Sub(current).Hours()/24)
}
