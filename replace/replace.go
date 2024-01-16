package replace

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var filter []string

// Do visit file
func doReplace(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err) // can't walk here,
		return nil
	}
	if info.IsDir() {
		return nil // not a file, ignore.
	}

	// To lower case checking
	lowPath := strings.ToLower(path)
	perm := info.Mode().Perm()

	for _, ft := range filter {
		if strings.HasSuffix(path, ft) || strings.HasSuffix(lowPath, ft) {

			// Let replace
			bData, err := os.ReadFile(path)
			if err != nil {
				fmt.Println(err)
			}

			// Convert to string
			sData := string(bData)

			// Loop to replace
			sData = loopReplaces(sData)

			// Write file back
			err = os.WriteFile(path, []byte(sData), perm)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Replaced at: " + path)
			}
		}
	}

	return nil
}

// Looping to replace
func loopReplaces(sData string) string {

	for _, rpItem := range listConfig {
		fo := rpItem.From
		to := rpItem.To
		if fo != "" {
			sData = strings.Replace(sData, fo, to, -1)
		}
	}

	return sData
}

// Replaces all strings in files
func Replaces(rootDir string, extensions []string) error {
	filter = extensions
	err := filepath.Walk(rootDir, doReplace)
	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", rootDir, err)
	}
	return err
}
