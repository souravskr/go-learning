package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go func() {
		for {
			msg := <-ch
			log.Print(msg)
		}
	}()

	fmt.Println("Type anything: ")

	i := 0
	for i <= 3 {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
		i++
	}
}
