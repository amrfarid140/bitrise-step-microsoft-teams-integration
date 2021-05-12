package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/colorstring"
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

func valueOptionToBool(userValue string) bool {
	return userValue == "yes"
}

func parseTimeString(cfg config) string {
	var timeAtLoc time.Time
	i, err := strconv.ParseInt(cfg.BuildTime, 10, 64)
	if err != nil {
		return string("Couldn't parse build time")
	}
	loc, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		fmt.Println(colorstring.Redf("\n%s", err))
		fmt.Println(colorstring.Cyan("\nExporting time in UTC...\n"))

		timeAtLoc = time.Unix(i, 0).In(time.UTC)
		return timeAtLoc.Format(time.RFC1123)
	}
	timeAtLoc = time.Unix(i, 0).In(loc)

	return timeAtLoc.Format(time.RFC1123)
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
	imagesSection := buildImagesSection(cfg)
	factsSection := buildFactsSection(cfg, buildSuccessful)
	message.Sections = []Section{primarySection, imagesSection, factsSection}

	// MessageCard Actions
	actions := []OpenURIAction{}
	if valueOptionToBool(cfg.EnableDefaultActions) {
		goToRepoAction := buildURIAction(Action{
			Text: "Go To Repo",
			Targets: []ActionTarget{
				{
					URI: cfg.RepoURL,
				},
			},
		})
		goToBuildAction := buildURIAction(Action{
			Text: "Go To Build",
			Targets: []ActionTarget{
				{
					URI: cfg.BuildURL,
				},
			},
		})
		actions = append(actions, goToRepoAction, goToBuildAction)
	}
	customActions := parseActions(cfg.Actions)
	for _, action := range customActions {
		actions = append(actions, buildURIAction(action))
	}
	message.Actions = actions

	return message
}

// Builds the primary section of the MessageCard content
func buildPrimarySection(cfg config) Section {
	section := Section{}
	section.ActivityTitle = cfg.SectionTitle
	section.ActivitySubtitle = cfg.SectionSubtitle
	section.Text = cfg.SectionText
	section.ActivityImage = cfg.SectionHeaderImage
	section.Markdown = valueOptionToBool(cfg.EnablePrimarySectionMarkdown)
	return section
}

// Builds a Section containing a list of Image
func buildImagesSection(cfg config) Section {
	section := Section{}
	if cfg.SectionImage != "" {
		image := Image{
			Image: cfg.SectionImage,
			Title: cfg.SectionImageDescription,
		}
		section.Images = []Image{image}
	}
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

	buildTimeFact := Fact{
		Name:  "Build Triggered",
		Value: parseTimeString(cfg),
	}

	workflowFact := Fact{
		Name:  "Workflow",
		Value: cfg.Workflow,
	}

	return Section{
		Markdown: valueOptionToBool(cfg.EnableBuildFactsMarkdown),
		Facts:    []Fact{buildStatusFact, buildNumberFact, buildBranchFact, buildTimeFact, workflowFact},
	}
}

func buildURIAction(action Action) OpenURIAction {
	uriAction := OpenURIAction{}
	uriAction.Type = "OpenUri"
	uriAction.Name = action.Text
	targets := []Target{}
	for _, target := range action.Targets {
		uriTarget := Target{}
		uriTarget.OS = optionalUserValue("default", target.OS)
		uriTarget.URI = target.URI
		targets = append(targets, uriTarget)
	}
	uriAction.Targets = targets

	return uriAction
}

// postMessage sends a message to a channel.
func postMessage(webhookURL string, msg Message, debugEnabled bool) error {
	b, err := json.MarshalIndent(msg, "", "  ")
	if err != nil {
		return err
	}
	if debugEnabled {
		log.Print(colorstring.Yellowf("\nRequest to Microsoft Teams:\n%s", b))
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
		log.Fatalf("Error: %s\n", err)
	}
	stepconf.Print(cfg)

	message := newMessage(cfg, buildSucceeded)
	if err := postMessage(cfg.WebhookURL, message, valueOptionToBool(cfg.EnableDebug)); err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println(colorstring.Cyan("\nMessage successfully sent!"))
}
