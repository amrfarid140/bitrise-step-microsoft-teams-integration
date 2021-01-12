package main

// config will be populated with the retrieved values from environment variables
// configured as step inputs.
type config struct {
	BuildNumber string `env:"BITRISE_BUILD_NUMBER"`
	AppTitle    string `env:"BITRISE_APP_TITLE"`
	AppURL      string `env:"BITRISE_APP_URL"`
	BuildURL    string `env:"BITRISE_BUILD_URL"`
	GitBranch   string `env:"BITRISE_GIT_BRANCH"`
	AppImageURL string `env:"BITRISE_APP_SLUG"`

	WebhookURL string `env:"webhook_url,required"`

	// Optional user inputs
	RepoURL string `env:"repository_url"`
}
