package main

import (
	"fmt"
	"time"
)

func main() {
	i := 10
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i - 1
	}
	fmt.Println("Countdown complete")
}
