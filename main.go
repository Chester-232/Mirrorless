package main

import (
	"fmt"

	"github.com/Chester-232/Mirrorless/client"
	"github.com/Chester-232/Mirrorless/server"
)

func main() {
	fmt.Println("Hello from the main file....");
	client.Test();
	server.Test();
}