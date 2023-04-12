package Log

import "log"

func ShowLog(ch chan string) {
	for {
		msg := <-ch
		log.Print(msg)
	}
}
