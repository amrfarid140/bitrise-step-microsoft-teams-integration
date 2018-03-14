package main

// Message to send to Microsoft Teams.
type Message struct {
	Type       string       `json:"@type"`
	Context    string       `json:"@context"`
	ThemeColor string       `json:"themeColor"`
	Summary    string       `json:"summary"`
	Markdown   string       `json:"markdown"`
	Sections   []Section    `json:"sections"`
	Actions    []ActionCard `json:"potentialAction"`
}

// Section to be shown in the message
type Section struct {
	AppTitle    string `json:"activityTitle"`
	BuildNumber string `json:"activitySubtitle"`
	AppImage    string `json:"activityImage"`
	Facts       []Fact `json:"facts"`
}

// Fact realted to the message
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ActionCard to be added to the message
type ActionCard struct {
	Type    string        `json:"@type"`
	Name    string        `json:"name"`
	Inputs  []ActionInput `json:"inputs"`
	Actions []Action      `json:"actions"`
}

// ActionInput for actions if any
type ActionInput struct {
	Type  string `json:"@type"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

// Action to be taken by the action card
type Action struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Target string `json:"target"`
}
