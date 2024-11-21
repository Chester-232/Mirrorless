package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Chester-232/Mirrorless/utils"
)

func Run() {
	var IP string;
	fmt.Print("Enter the server IP: ")
	fmt.Scan(&IP);
	conn, err := net.Dial("tcp", IP+ ":8080");
	utils.Checker("Error connecting to the server: ", err);
	defer conn.Close()

	fmt.Println("Connected to server.");
	destDir := "./destination";
	os.MkdirAll(destDir, 0755);

	for {
		reader := bufio.NewReader(conn);

		meta, err := reader.ReadString('\n');
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by the server.");
				return;
			}
			fmt.Println("Error reading metadata: ", err);
			return;
		}
		meta = strings.TrimSpace(meta);
		parts := strings.Split(meta, "|");
		if len(parts) != 2 {
			fmt.Println("Invalid metadata format.")
			continue;
		}
		fileName := parts[0];
		fileSize, _ := strconv.ParseInt(parts[1], 10, 64);

		destPath := filepath.Join(destDir, fileName);
		outFile, err := os.Create(destPath)
		if err != nil {
			fmt.Println("Error creating file: ", err);
			return;
		}
		written, err := io.CopyN(outFile, reader, fileSize);
		outFile.Close()
		if err != nil && err != io.EOF {
			fmt.Println("Error writing File: ", err)
		}
		fmt.Printf("Received file %s (%d bytes)\n", fileName, written);
	}
}