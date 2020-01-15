# Changelog
All notable changes to this project will be documented in this file.

The format of this file is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), 
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] 

## [1.0.2] - 2020-01-15
### Changed
- Modified release.sh so that it now updates the version in Dockerfile's LABEL step
- Changed how HTTP 4xx and 5xx bodies are read and printed in the logger

## [1.0.1] - 2020-01-14
### Added
- Added the changelog.md file 
### Changed
- Modified version of release.sh to manage changelog.md and README.md

## [N/A] - 2020-01-14
### Changed
- Renaming from 'jira-api-resource' to 'jira-api-issue-resource'

## [1.0.0] - 2020-01-08
Initial version that supports
* Read issues
* Update issues
* 'force-on-parent' mechanic
