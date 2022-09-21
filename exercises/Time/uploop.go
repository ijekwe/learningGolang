package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i > 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i + 1
		if i == 10 {
			fmt.Println("Complete", i)
			break
		}
	}
}
