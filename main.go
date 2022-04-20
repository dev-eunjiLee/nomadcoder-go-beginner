package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	go sexyCount("flynn")
	time.Sleep(time.Second * 2)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second) // * 1초 슬립
	}
}
