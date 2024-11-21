package main

import (
	"fmt"

	"github.com/Chester-232/Mirrorless/client"
	"github.com/Chester-232/Mirrorless/server"
)

func main() {
	var choice int;
	
	mainloop:
	for {
		fmt.Print("Which side are you running?\n1. Server\n2.Client\n3.Quit\n->");
		fmt.Scan(&choice);
		switch choice {
		case 1:
			server.Run()
		case 2:
			client.Run()
		case 3:
			break mainloop;
		default:
			fmt.Println("Invalid choice.")
		}
	}
}