package main

// config will be populated with the retrieved values from environment variables
// configured as step inputs.
type config struct {
	// Bitrise environment inputs
	BuildNumber       string `env:"BITRISE_BUILD_NUMBER"`
	AppTitle          string `env:"BITRISE_APP_TITLE"`
	AppURL            string `env:"BITRISE_APP_URL"`
	BuildURL          string `env:"BITRISE_BUILD_URL"`
	BuildTime         string `env:"BITRISE_BUILD_TRIGGER_TIMESTAMP"`
	CommitMessageBody string `env:"GIT_CLONE_COMMIT_MESSAGE_BODY"`
	GitBranch         string `env:"BITRISE_GIT_BRANCH"`
	Workflow          string `env:"BITRISE_TRIGGERED_WORKFLOW_TITLE"`

	// Required user inputs
	WebhookURL string `env:"webhook_url,required"`

	// Optional user inputs
	CardTitle                    string `env:"card_title"`
	SuccessThemeColor            string `env:"success_theme_color"`
	FailedThemeColor             string `env:"failed_theme_color"`
	SectionTitle                 string `env:"section_title"`
	SectionSubtitle              string `env:"section_subtitle"`
	SectionText                  string `env:"section_text"`
	SectionHeaderImage           string `env:"section_header_image"`
	SectionImage                 string `env:"section_image"`
	SectionImageDescription      string `env:"section_image_description"`
	EnablePrimarySectionMarkdown string `env:"enable_primary_section_markdown"`
	EnableBuildFactsMarkdown     string `env:"enable_build_status_facts_markdown"`
	EnableDefaultActions         string `env:"enable_default_actions"`
	EnableDebug                  string `env:"enable_debug"`
	RepoURL                      string `env:"repository_url"`
	Actions                      string `env:"actions"`
	Timezone                     string `env:"timezone"`
}
