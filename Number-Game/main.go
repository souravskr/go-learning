package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and ENTER when you are ready."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	// Instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 to 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought off", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// Answer
	answer = firstNumber*secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
