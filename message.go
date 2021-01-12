package main

// Message to send to Microsoft Teams.
type Message struct {
	Type       string          `json:"@type"`
	Context    string          `json:"@context"`
	ThemeColor string          `json:"themeColor"`
	Summary    string          `json:"summary"`
	Markdown   string          `json:"markdown"`
	Sections   []Section       `json:"sections"`
	Actions    []OpenUriAction `json:"potentialAction"`
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

// OpenUri action for link buttons
type OpenUriAction struct {
	Type    string   `json:"@type"`
	Name    string   `json:"name"`
	Targets []Target `json:"targets"`
}

// The required Target object that resides inside `OpenUriAction`s
type Target struct {
	OS  string `json:"os"`
	Uri string `json:"uri"`
}
