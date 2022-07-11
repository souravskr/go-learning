package main

import (
	"Go-Learning/doctor"
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	var input string
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for {
		//userInput, _ := reader.ReadString('\n')
		fmt.Scanln(&input)
		if input == "quit" {
			break
		}
		fmt.Println(input)
		response := doctor.Response(input)
		fmt.Println(response)
	}
}
