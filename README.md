# Jira Helper CLI
This project is born to help me and my colleague @cesto93 to automate the workflow we had with Jira in our project. This CLI is a 1:1 integration with the [Jira REST API v2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/#about) to help you execute repetitive tasks via scripts or simply change your card status, assignee and more to come.

## Featrues
Because this implementation aims to be 1:1 implementation of the [JIRA REST API v2](https://developer.atlassian.com/cloud/jira/platform/rest/v2/intro/#about) a command will exist for each section. Each command will have it's own subcommand, each subcommand will be linked to a URL and an HTTP method to keep the integration simple.

To follow you can find the implemented endpoints for each jira group:
- [Issues](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-group-issues)
    - [Assign Issue](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-assignee-put) 🔜
    - [Get Transitions](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-get) ✅
    - [Transition Issue](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-post) ✅

- [Issue search](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-group-issue-search)
    - [Search for Issue GET](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-get) ❌ using directly the POST version

    - [Search for Issue POST](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-get) ✅

- [Myself](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-group-myself)
    - [Get Current User](https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-myself/#api-rest-api-2-myself-get) 🔜

If any body is required a `--body` flag is provided and can be added there.
If any query or path parameter is required then a dedicated flag is added for the command and must be filled if required.

## Configuration
Because it is possible to use multiple accounts a `toml` configuration file is provided with in order to select a profile with the `--profile` flag when launching a command.
The configuration file will be placed in `~/.jhelp.config`
