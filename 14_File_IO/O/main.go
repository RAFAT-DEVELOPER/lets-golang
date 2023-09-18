package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	text := "Hello, file I/O!"
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully")
}
