package client

import (
	"net"

	"github.com/Chester-232/Mirrorless/utils"
)

func main() {
	conn, err := net.Dial("tcp", ":8080");
	utils.Checker("Error connecting to the server: ", err);
	defer conn.Close()
}