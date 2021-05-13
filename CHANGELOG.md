## [2.1.1] - 13th May 2021

* [#32] Fix `step.yml` indentation

## [2.1.0] - 13th May 2021

* [#28] Allow setting a custom timezone for 'Build Triggered' time
* [#29] Added ability to add custom image to `MessageCard`. 

## [2.0.0] - 28th Feb 2021

* [#14] Added ability to declare custom `MessageCard` actions via JSON string
* [#17] Bumped step to use Go 1.16
* [#18] Added additional Go test steps to `step.yml`
* [#19] Fix `MessageCard` `ActivityImage` input
	* now correctly displays an image if image URL declared in step input
* [#20] BREAKING: Step now uses "yes|no" values for inputs that were previously `bool` values


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
