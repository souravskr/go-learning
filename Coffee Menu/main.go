package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"strconv"
)

func main() {
	_ = keyboard.Open()
	defer func() {
		_ = keyboard.Close()
	}()

	menu := map[int]string{1: "Cappuccino", 2: "Latte", 3: "Americano", 4: "Mocha", 5: "Macchiato", 6: "Espresso"}

	fmt.Println("MENU")
	fmt.Println("-------")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")
	fmt.Println("========================")

	char := ' '
	for char != 'q' {
		char, _, _ = keyboard.GetSingleKey()
		option, _ := strconv.Atoi(string(char))
		fmt.Println("->", option)
		fmt.Printf("You chose %s\n", menu[option])
	}
	fmt.Println("Program quiting...")
}
