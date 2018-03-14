package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitrise-tools/go-steputils/stepconf"
)

func newMessage(buildOk bool, cfg config) Message {
	message := Message{}
	message.Type = "MessageCard"
	message.Context = "http://schema.org/extensions"
	message.ThemeColor = "0076D7"
	if buildOk {
		message.Summary = fmt.Sprintf("%s build number %s succeeded", cfg.AppTitle, cfg.BuildNumber)
	} else {
		message.Summary = fmt.Sprintf("%s build number %s failed", cfg.AppTitle, cfg.BuildNumber)
	}
	message.Markdown = "true"
	section := Section{}
	section.AppTitle = fmt.Sprintf("![AppImage](%s)%s", cfg.AppImageURL, cfg.AppTitle)
	section.BuildNumber = cfg.BuildNumber
	section.AppImage = cfg.AppImageURL

	buildStatusFact := Fact{}
	buildStatusFact.Name = "Build Status"
	if buildOk {
		buildStatusFact.Value = `<span style="color:green">Success</span>`
	} else {
		buildStatusFact.Value = `<span style="color:red">Fail</span>`
	}

	buildNumberFact := Fact{}
	buildNumberFact.Name = "Build Number"
	buildNumberFact.Value = cfg.BuildNumber

	buildBranchFact := Fact{}
	buildBranchFact.Name = "Git Branch"
	buildBranchFact.Value = cfg.GitBranch

	section.Facts = []Fact{buildStatusFact, buildNumberFact, buildBranchFact}

	message.Sections = []Section{section}

	goToRepoActionCard := ActionCard{}
	goToRepoActionCard.Type = "ActionCard"
	goToRepoActionCard.Name = "Go To Repo"
	goToRepoActionCardAction := Action{}
	goToRepoActionCardAction.Type = "HttpGet"
	goToRepoActionCardAction.Name = "Go To Repo"
	goToRepoActionCardAction.Target = cfg.RepoURL

	goToBuildActionCard := ActionCard{}
	goToBuildActionCard.Type = "ActionCard"
	goToBuildActionCard.Name = "Go To Build"
	goToBuildActionCardAction := Action{}
	goToBuildActionCardAction.Type = "HttpGet"
	goToBuildActionCardAction.Name = "Go To Build"
	goToBuildActionCardAction.Target = cfg.BuildURL

	message.Actions = []ActionCard{goToRepoActionCard, goToBuildActionCard}

	return message
}

// postMessage sends a message to a channel.
func postMessage(webhookURL string, msg Message) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Request to Microsoft Teams: %s", webhookURL))

	resp, err := http.Post(webhookURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("failed to send the request: %s", err)
	}
	defer func() {
		if cerr := resp.Body.Close(); err == nil {
			err = cerr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("server error: %s, failed to read response: %s", resp.Status, err)
		}
		return fmt.Errorf("server error: %s, response: %s", resp.Status, body)
	}

	return nil
}

func main() {
	var cfg config
	if err := stepconf.Parse(&cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	stepconf.Print(cfg)

	// success is true if the build is successful, false otherwise.
	success := os.Getenv("BITRISE_BUILD_STATUS") == "0"
	message := newMessage(success, cfg)
	if err := postMessage(cfg.WebhookURL, message); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	fmt.Println("Message successfully sent!")
}
