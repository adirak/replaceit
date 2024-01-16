package replace

import (
	"encoding/json"
	"os"
)

const replaceConfPath = "replaces.json"

var listConfig []ReplaceItem

// Read Config file
func init() {

	bData, err := os.ReadFile(replaceConfPath)
	if err != nil {
		panic(err)
	}

	// Init
	listConfig = []ReplaceItem{}

	// Read Replace files
	err = json.Unmarshal(bData, &listConfig)
	if err != nil {
		panic(err)
	}

}
