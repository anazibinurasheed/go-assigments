package main

import (
	"fmt"
	"time"
)

func main() {
	PresentTime := time.Now()
	fmt.Println(PresentTime)
	fmt.Println(PresentTime.Format(" 01-02-2006 Monday 15:04:05"))          // IN this  way we are setup  formating the time , this numerals used hear
	CreatedDate := time.Date(2003, time.August, 5, 23, 59, 45, 0, time.UTC) // is the only numerals that change the format
	fmt.Println(CreatedDate)
	
}