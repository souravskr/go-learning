package main

import (
	doctor "Variables/Doctor"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	whatToSay := doctor.Intro()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(whatToSay)
	for {
		fmt.Print("> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if strings.ToLower(userInput) == "quit" {
			fmt.Println("Bye...")
			break
		} else {
			fmt.Println()
			fmt.Println(doctor.Response(userInput))
		}

	}
}
