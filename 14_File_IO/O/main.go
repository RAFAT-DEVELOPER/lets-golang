package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Define a file path
	filePath := "example.txt"

	// Case 1: Writing to a File
	data := []byte("Hello, this is a sample text for file I/O in Go.")
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully.")

	// Case 2: Reading from a File
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("File content:", string(content))

	// Case 3: Checking File Existence
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		} else {
			fmt.Println("Error checking file existence:", err)
		}
	} else {
		fmt.Println("File exists. File size:", fileInfo.Size(), "bytes")
	}

	// Case 4: Creating a Directory
	dirPath := "mydir"
	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created:", dirPath)

	// Case 5: Removing a File
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed:", filePath)

	// Case 6: Removing a Directory
	err = os.Remove(dirPath)
	if err != nil {
		fmt.Println("Error removing directory:", err)
		return
	}
	fmt.Println("Directory removed:", dirPath)
}
