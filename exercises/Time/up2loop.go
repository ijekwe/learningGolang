package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i >= 0 {
		fmt.Println(i)
		time.Sleep(time.Second)
		i = i + 2
		if i == 10 {
			fmt.Println(i)
			break
		}
	}
}
