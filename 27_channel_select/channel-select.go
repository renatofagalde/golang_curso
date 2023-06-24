package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "canal 01"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "canal 02"
		}
	}()

	for {
		select {
		case msg01 := <-canal1:
			fmt.Println(msg01)
		case msg02 := <-canal2:
			fmt.Println(msg02)
		}
	}

}
