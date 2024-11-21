package server

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/Chester-232/Mirrorless/utils"
	"github.com/fsnotify/fsnotify"
)

func Run() {
	watcher, err := fsnotify.NewWatcher()
	utils.Checker("Error creating the watcher: ", err)
	defer watcher.Close();

	sourceDir := "server Folder/";
	err = watcher.Add(sourceDir);
	utils.Checker("Error watching the dir: ", err);

	listener, err := net.Listen("tcp", ":8080");
	utils.Checker("Error starting server: ", err);
	defer listener.Close();

	fmt.Println("Server started. waiting for clinets...");

	conn, err := listener.Accept();
	utils.Checker("Error accepting clients: ", err);
	defer conn.Close();
	fmt.Println("Client Connected.");
	for {
		select {
		case event := <-watcher.Events:
			fmt.Println("Event: ", event)
			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
				filePath := event.Name;
				err := sendFile(filePath, conn);
				utils.Checker("Error sending file: ", err);
			}
		case err := <-watcher.Errors:
			fmt.Println("Watcher Eror: ", err);
		}
	}
}

func sendFile(filePath string, conn net.Conn) error {
	file, err := os.Open(filePath);
	if err != nil { return err }
	defer file.Close()

	fileInfo, _ := file.Stat();
	conn.Write([]byte(fmt.Sprintf("%s|%d\n", fileInfo.Name(), fileInfo.Size())))

	_, err = io.Copy(conn, file);
	return err
}