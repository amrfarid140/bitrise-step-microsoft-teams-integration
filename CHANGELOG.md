## [1.2.2] - 18th Feb 2021

* Clean up Gopkg.* files, no longer required since migrating to Go Modules
* Remove deprecated bitrise-tools dependency in favour of bitrise-io
* Removed unnecessary Go brew deps from `step.yml`
	* Only go toolkit is required

## [1.2.1] - 6th Feb 2021

* fix summaries within `step.yml`

## [1.2.0] - 24th Jan 2021

* MessageCard customisation
	* ability to override each text field within the card
    * defaults to Bitrise environment variables if custom values not set
* Migration to Go Modules
* Unit tests and GitHub Actions CI checks

## [1.1.0] - 13th Jan 2021

* Added new user input to override `$GIT_REPOSITORY_URL` for SSH Bitrise projects.

## [1.0.2] 12th Jan 2021

* Fixed URI buttons not appearing in Teams card
* Fixed CHANGELOG.md typo and created release notes
