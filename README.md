# Jira Helper CLI
This project is born to help building automation scripts with jira. This CLI is a 1:1 integration with the [Jira REST API v2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/#about) to help you execute repetitive tasks via scripts or simply change your card status, assignee and more to come.

Because a tool to query JSON payloads is already available ([jq](https://jqlang.github.io/jq/)) this cli is intended to act as a simple HTTP client so it's flexible enogh and can help building scripts to automate you work with jira if needed.

## Installing
```
    go install github.com/cedrata/jira-helper@latest
```

## Featrues
Because this implementation aims to be 1:1 implementation of the [JIRA REST API v2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/#about) a command will exist for each section. Each command will have it's own subcommand, each subcommand will be linked to a URL and an HTTP method to keep the integration simple.

To follow you can find the implemented endpoints for each jira group:
- [Issues](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-group-issues)
    - [Assign Issue](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-assignee-put) ✅
    - [Get Transitions](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-get) ✅
    - [Transition Issue](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-post) ✅

- [Issue search](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-group-issue-search)
    - [Search for Issue GET](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-get) ❌ using directly the POST version

    - [Search for Issue POST](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-get) ✅

- [Myself](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-group-myself)
    - [Get Current User](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-rest-api-2-myself-get) ✅

## Configuration
Before using the cli create the file `~/.jira-helper.config` with the same structure as follows:
```
[default]
host=your.host.com
token=yourjiratoken
```

The configuration file is strucutred to hava many profiles. The profile `default` is the one used if the `--profile` flag is not provided. If you wish to have more than one profile and call it explicitely add a section structured like the example before and you can add the `--profile` flag with the name you set for the profile.
