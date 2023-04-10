package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

var chanKey chan rune

func main() {
	chanKey = make(chan rune)
	go listenForKeyPress()

	_ = keyboard.Open()
	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Press any key.")
	for {
		key, _, _ := keyboard.GetSingleKey()
		if key == 'Q' || key == 'q' {
			fmt.Println("Quiting...")
			break
		}
		chanKey <- key
	}
}

func listenForKeyPress() {
	for {
		key := <-chanKey
		fmt.Println("You pressed", key)
	}
}
