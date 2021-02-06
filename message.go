package main

// Message to send to Microsoft Teams.
type Message struct {
	Type       string          `json:"@type"`
	Context    string          `json:"@context"`
	ThemeColor string          `json:"themeColor"`
	Title      string          `json:"title"`
	Summary    string          `json:"summary"`
	Sections   []Section       `json:"sections"`
	Actions    []OpenURIAction `json:"potentialAction"`
}

// Section to be shown in the message
type Section struct {
	ActivityTitle    string    `json:"activityTitle"`
	ActivitySubtitle string    `json:"activitySubtitle"`
	ActivityImage    string    `json:"activityImage"`
	Facts            []Fact    `json:"facts"`
	Markdown         bool      `json:"markdown"`
	Text             string    `json:"text"`
	HeroImage        HeroImage `json:"heroImage"`
}

// HeroImage that is displayed within the Message
type HeroImage struct {
	Image string `json:"image"`
}

// Fact related to the message
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// OpenURIAction action for link buttons
type OpenURIAction struct {
	Type    string   `json:"@type"`
	Name    string   `json:"name"`
	Targets []Target `json:"targets"`
}

// Target object that resides inside `OpenUriAction`s
type Target struct {
	OS  string `json:"os"`
	URI string `json:"uri"`
}
