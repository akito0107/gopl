package main

import "fmt"

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.board <- fmt.Sprintf("%s\n", msg)
			}
		case cli := <-entering:
			for c := range clients {
				cli.board <- fmt.Sprintf("username: %s is in room", c.name)
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.board)
		}
	}
}
