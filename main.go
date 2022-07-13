package main

import (
	"Go-Learning/doctor"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for {
		fmt.Print("--> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		}
		fmt.Println(doctor.Response(userInput))
	}
}
