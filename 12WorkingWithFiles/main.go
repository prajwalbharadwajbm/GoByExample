package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createFile(filename string) error {
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error Creating a file: ", err)
	}
	return err
}

func readData(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Reading the file: ", err)
	}
	return string(data)
}

func writeData(filename string, data string) {
	err := os.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		fmt.Println("Error Writing to the file: ", err)
	}
	fmt.Println("Successfully Wrote to file: ", filename)
}

func returnFilePath(filename string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	var resultPath string
	err = filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == filename {
			resultPath = path
			return fmt.Errorf("file found") // Stop walking after finding the file
		}
		return nil
	})
	if err != nil && err.Error() != "file found" {
		return "", err
	}
	return resultPath, nil
}

func main() {
	fmt.Println("Welcome to Filecreator")
	fmt.Println()
	fmt.Println("Please enter,")
	fmt.Print("Filename: ")
	scanner := bufio.NewReader(os.Stdin)
	filename, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Couldn't read file from stdin: ", err)
	}
	filename = strings.TrimSpace(filename)
	fmt.Print("Data to be added to the file: ")
	data, _, err := scanner.ReadLine()
	if err != nil {
		fmt.Println("Couldn't read data from stdin: ", err)
	}
	var filePath string
	createErr := createFile(string(filename))
	if createErr == nil {
		var filePathErr error
		filePath, filePathErr = returnFilePath(string(filename))
		if filePathErr != nil {
			fmt.Println("Couldn't find the file: ", err)
		}
	}
	writeData(string(filename), string(data))
	fmt.Printf("Your file has been created in the path %s, please do validate if it worked!! :)\n", filePath)
	fmt.Println("\nYour File Contents: ", readData(filename))
}
