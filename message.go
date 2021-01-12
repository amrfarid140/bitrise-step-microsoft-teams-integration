package main

// Message to send to Microsoft Teams.
type Message struct {
	Type       string          `json:"@type"`
	Context    string          `json:"@context"`
	ThemeColor string          `json:"themeColor"`
	Title      string          `json:"title"`
	Summary    string          `json:"summary"`
	Markdown   string          `json:"markdown"`
	Sections   []Section       `json:"sections"`
	Actions    []OpenUriAction `json:"potentialAction"`
}

// Section to be shown in the message
type Section struct {
	ActivityTitle    string    `json:"activityTitle"`
	ActivitySubtitle string    `json:"activitySubtitle"`
	ActivityImage    string    `json:"activityImage"`
	Facts            []Fact    `json:"facts"`
	Text             string    `json:"text"`
	HeroImage        HeroImage `json:"heroImage"`
}

type HeroImage struct {
	Image string `json:"image"`
}

// Fact related to the message
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
