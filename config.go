package main

// Config will be populated with the retrieved values from environment variables
// configured as step inputs.
type Config struct {

	// Message
	BuildNumber string `env:"BITRISE_BUILD_NUMBER"`
	AppTitle    string `env:"BITRISE_APP_TITLE"`
	AppURL      string `env:"BITRISE_APP_URL"`
	BuildURL    string `env:"BITRISE_BUILD_URL"`
	RepoURL     string `env:"GIT_REPOSITORY_URL"`
	GitBranch   string `env:"BITRISE_GIT_BRANCH"`
	AppImageURL string `env:"BITRISE_APP_SLUG"`
}
