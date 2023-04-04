package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	_ = keyboard.Open()
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		key, _, _ := keyboard.GetSingleKey()
		if key == 'q' || key == 'Q' {
			fmt.Println("Quiting...")
			break
		}
		keyPressChan <- key
	}
}
