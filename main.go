package main

import (
	"adirak/replaceit/replace"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: replaceit <root_dir> <extension1> <extension2> ...")
		os.Exit(1)
	}

	fmt.Println("root_dir: " + os.Args[1])

	extensions := []string{}
	for i := 2; i < len(os.Args); i++ {
		ext := strings.TrimSpace(os.Args[i])
		extensions = append(extensions, ext)
	}
	fmt.Printf("extensions: %v \n", extensions)

	// Call Replace function
	replace.Replaces(os.Args[1], extensions)

}
