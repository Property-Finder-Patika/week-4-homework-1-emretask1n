package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//WalkDir Func
	var dirPath string
	fmt.Println("\nWelcome to the List Files program")
	fmt.Println("Example for path to list of files => C:\\Users\\e-tas\\Desktop\\week-4-homework-1-emretask1n")
	fmt.Println("Please enter a valid directory:")
	_, err := fmt.Scanf("%s", &dirPath)
	if err != nil {
		fmt.Println("That is not a valid directory!")
		os.Exit(0)
	}
	err2 := filepath.WalkDir(dirPath,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			return nil
		})
	if err2 != nil {
		log.Println(err2)
	}
}
