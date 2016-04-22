package main

import "fmt"
import "github.com/andrewarrow/paradise_ftp/server"

func main() {
	fmt.Println("Using paradise_ftp")
	go server.Monitor()
	fm := NewGitHubFileManager()
	server.Start(fm)
}
