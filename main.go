package main

import (
	"Go-Learning/doctor"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	userInput, _ := reader.ReadString('\n')
	fmt.Println(userInput)
}
