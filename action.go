package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitrise-io/go-utils/colorstring"
)

// Action object within the `actions` JSON object array.
type Action struct {
	Text    string `json:"text"`
	Targets []ActionTarget
}

// ActionTarget which defines the platform in which the Action should be available
type ActionTarget struct {
	OS  string `json:"os"`
	URI string `json:"uri"`
}

func parseActions(jsonString string) []Action {
	var actionList []Action
	err := json.Unmarshal([]byte(jsonString), &actionList)
	if err != nil {
		fmt.Println(colorstring.Redf("Couldn't Unmarshal JSON: %v, \n %s", jsonString, err))
	}
	fmt.Printf("JSON value: %v", actionList)
	return actionList
}
