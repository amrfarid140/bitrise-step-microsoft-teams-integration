package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bitrise-io/go-steputils/stepconf"
)

var buildSucceeded = os.Getenv("BITRISE_BUILD_STATUS") == "0"

func getValueForBuildStatus(ifSuccess, ifFailed string, buildSucceeded bool) string {
	if buildSucceeded || ifFailed == "" {
		return ifSuccess
	}
	return ifFailed
}

func optionalUserValue(defaultValue, userValue string) string {
	if userValue == "" {
		return defaultValue
	}
	return userValue
}

func newMessage(cfg config, buildSuccessful bool) Message {
	message := Message{}
	message.Type = "MessageCard"
	message.Context = "http://schema.org/extensions"
	message.ThemeColor = getValueForBuildStatus(
		cfg.SuccessThemeColor,
		cfg.FailedThemeColor,
		buildSuccessful,
	)
	message.Title = optionalUserValue(cfg.AppTitle, cfg.CardTitle)
	message.Summary = getValueForBuildStatus(
		fmt.Sprintf("%s #%s succeeded", cfg.AppTitle, cfg.BuildNumber),
		fmt.Sprintf("%s #%s failed", cfg.AppTitle, cfg.BuildNumber),
		buildSuccessful,
	)

	// MessageCard sections
	primarySection := buildPrimarySection(cfg)
	factsSection := buildFactsSection(cfg, buildSuccessful)
	message.Sections = []Section{primarySection, factsSection}

	// MessageCard Actions
	goToRepoAction := buildUriButton("Go To Repo", cfg.RepoURL)
	goToBuildAction := buildUriButton("Go To Build", cfg.BuildURL)
	message.Actions = []OpenUriAction{goToRepoAction, goToBuildAction}

	return message
}

// Builds the primary section of the MessageCard content
func buildPrimarySection(cfg config) Section {
	section := Section{}
	section.ActivityTitle = cfg.SectionTitle
	section.ActivitySubtitle = cfg.SectionSubtitle
	section.Text = cfg.SectionText
	// TODO: fix AppImageURL
	section.ActivityImage = cfg.AppImageURL
	section.Markdown = cfg.EnablePrimarySectionMarkdown
	return section
}

// Builds a Section containing a list of Fact related to build status
func buildFactsSection(cfg config, buildSuccessful bool) Section {
	buildStatusFact := Fact{
		Name: "Build Status",
		Value: getValueForBuildStatus(
			fmt.Sprintf(`<span style="color:#%s">Success</span>`, cfg.SuccessThemeColor),
			fmt.Sprintf(`<span style="color:#%s">Fail</span>`, cfg.FailedThemeColor),
			buildSuccessful,
		),
	}

	buildNumberFact := Fact{
		Name:  "Build Number",
		Value: cfg.BuildNumber,
	}

	buildBranchFact := Fact{
		Name:  "Git Branch",
		Value: cfg.GitBranch,
	}

	i, err := strconv.ParseInt(cfg.BuildTime, 10, 64)
	if err != nil {
		_ = fmt.Errorf("failed to parse the given build time: %s", err)
	}
	// Force UTC, as it otherwise defaults to locale of executing system
	parsedTime := time.Unix(i, 0).In(time.UTC)
	buildTimeFact := Fact{
		Name:  "Build Triggered",
		Value: parsedTime.Format(time.RFC1123),
	}

	workflowFact := Fact{
		Name:  "Workflow",
		Value: cfg.Workflow,
	}

	return Section{
		Markdown: cfg.EnableBuildFactsMarkdown,
		Facts:    []Fact{buildStatusFact, buildNumberFact, buildBranchFact, buildTimeFact, workflowFact},
	}
}

func buildUriButton(buttonText string, uri string) OpenUriAction {
	uriAction := OpenUriAction{}
	uriAction.Type = "OpenUri"
	uriAction.Name = buttonText
	uriTarget := Target{}
	uriTarget.OS = "default"
	uriTarget.Uri = uri
	uriAction.Targets = []Target{uriTarget}

	return uriAction
}

// postMessage sends a message to a channel.
func postMessage(webhookURL string, msg Message, debugEnabled bool) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Request to Microsoft Teams: %s", webhookURL))
	if debugEnabled {
		fmt.Println(fmt.Sprintf("JSON body: %s\n", b))
	}

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

	message := newMessage(cfg, buildSucceeded)
	if err := postMessage(cfg.WebhookURL, message, cfg.EnableDebug); err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err))
		os.Exit(1)
	}

	fmt.Println("Message successfully sent!")
}
