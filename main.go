package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}
	//result := <-c // 메인에서 채널의 결과를 기다리고 있으면, 메인 함수가 끝나지 않는다.
	//fmt.Println(result)

	//time.Sleep(time.Second * 3)
	for i := 0; i < len(people); i++ {

		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is Sexy"
}
