package main

// Message to post to a slack channel.
// See also: https://api.slack.com/methods/chat.postMessage
type Message struct {
	// Channel to send message to.
	//
	// Can be an encoded ID (eg. C024BE91L), or the channel's name (eg. #general).
	Type string `json:"@type"`

	// Text of the message to send. Required, unless providing only attachments instead.
	Context string `json:"@context"`

	ThemeColor string `json:"themeColor"`

	Summary string `json:"summary"`

	Markdown string `json:"markdown"`

	Sections []Section `json:"sections"`

	Actions []ActionCard `json:"potentialAction"`
}

//Section to be shown in the message
type Section struct {
	AppTitle    string `json:"activityTitle"`
	BuildNumber string `json:"activitySubtitle"`
	AppImage    string `json:"activityImage"`
	Facts       []Fact `json:"facts"`
}

//Fact realted to the message
type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//ActionCard to be added to the message
type ActionCard struct {
	Type    string        `json:"@type"`
	Name    string        `json:"name"`
	Inputs  []ActionInput `json:"inputs"`
	Actions []Action      `json:"actions"`
}

//ActionInput for actions if any
type ActionInput struct {
	Type  string `json:"@type"`
	ID    string `json:"id"`
	Title string `json:"title"`
}

//Action to be taken by the action card
type Action struct {
	Type   string `json:"@type"`
	Name   string `json:"name"`
	Target string `json:"target"`
}
