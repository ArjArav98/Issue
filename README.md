# Issue
It can sometimes be _frustrating as hell_ to switch to your browser from your IDE or terminal to view task requirements. It can be even more irritating when these websites don't have proper keyboard support.

To combat all this, **Issue** is a simple CLI utility for viewing and marking GitLab project issues _(Support for Jira, Github and Clubhouse coming soon!)_.


* [Quickstart](https://github.com/ArjArav98/Issue/blob/master/README.md#installation)
  * [Initial Setup](https://github.com/ArjArav98/Issue/blob/master/README.md#initial-setup)
* [Usage](https://github.com/ArjArav98/Issue/blob/master/README.md#usage)
   * [Listing Issues](https://github.com/ArjArav98/Issue/blob/master/README.md#usage)
   * [Viewing Single Issues](https://github.com/ArjArav98/Issue/blob/master/README.md#usage)
   * [Miscellaneous](https://github.com/ArjArav98/Issue/blob/master/README.md#usage)
   * [Listings Search Parameters](https://github.com/ArjArav98/Issue/blob/master/README.md#usage)
* [Potential Contributions](https://github.com/ArjArav98/Issue/blob/master/README.md#potential-contributions)
* [Tools Used](https://github.com/ArjArav98/Issue/blob/master/README.md#tools-used)

# Quickstart
* Download the executable corresponding to your OS from the table below.

   | OS & Architecture | Link |
   |---|---|
   |Darwin (Mac OS) AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/darwin_amd64/issue) |
   |Darwin (Mac OS) ARM64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/darwin_arm64/issue) |
   |Linux 386| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_386/issue) |
   |Linux AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_amd64/issue) |
   |Linux ARM| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_arm/issue) |
   |Linux ARM64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_arm64/issue) |
   |Windows 386| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/windows_386/issue.exe) |
   |Windows AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/windows_amd64/issue.exe) |

* Add the path of the directory, in which the executable is present, to your `PATH` environment variable.
* Run `issue version` to verify the successful installation!

### Initial Setup
* Run `issue init`.
* Open the generated `issues.config.json` file.
  * [Generate a Gitlab API token.](https://gitlab.com/-/profile/personal_access_tokens) Only provide read_* permissions to the token. Paste your token as the value for the `BearerToken` attribute in the config file.
  * The value for the `HostUrl` attribute must be the full domain URL at which Gitlab is hosted for you (ex; `https://mycompany.com/gitlab` or `https://gitlab.mycompany.com`). If you are using the cloud (normal) version of Gitlab, use `https://gitlab.com`.
  * The value for the `RepositoryNamespace` attribute must be the repository name (ex; ArjArav98/Issue or stark-industries/ironmansuit).

# Usage
|Issue Listings Commands|Description|
|---|---|
|`issue list` | Lists all issues.|
|`issue list --assignee_username darth.vader69 --labels "Doing,Backend"` | Lists all issues after applying search parameters. For full list of search parameters and examples, click here.|
|`issue list --my-open-issues` | Shortcut which displays all issues assigned to you, which are open.|
|`issue list --my-issues` | Shortcut which displays all issues assigned to you.|

|Detailed Issue View Commands|Description|
|---|---|
|`issue show <issue_id>` | Displays the selected issue in detail.|
|`issue show --with-comments <issue_id>` | Displays the selected issue in detail, along with its comments.|
|`issue show --no-comments <issue_id>` | Displays the selected issue's comments only.|

|Miscellaneous Commands|Description|
|---|---|
|`issue init` | Creates an empty config file in the current directory.|
|`issue version` | Displays the current version of the tool.|
|`issue help` | Displays a help menu with usage instructions for each command.|

---------
### Listings Search Parameters

|Search Parameter|Possible Values|
|---|---|
|`--assignee_id`| (integer/Any/None)|
|`--assignee_username`| (comma-separated-strings)|
|`--created_after`| (datetime)|
|`--created_before`| (datetime)|
|`--updated_after`| (datetime)|
|`--updated_before`| (datetime)|
|`--labels`| (comma-separated-strings)|
|`--search`| (string)|
|`--order_by`| (created_at/updated_at/)|
|`--state`| (opened/closed)|


**Examples;**
```
issue list --my-open-issues --labels backend,doing --created_before --created_before 2012-12-21
issue list --assignee_username sauron123 --assignee_username frodo99
```

# Troubleshooting & Common Errors
* If you encounter a 'permission denied' message on Mac OS, you might have to change the permissions of the executable file to 711. Once this is over, you will have to go to `System Preferences -> Security & Privacy` and allow Issue to be 'opened anyway'.

# Potential Contributions
Feel free to submit an MR for a feature addition or bug.

# Tools Used
* Gitlab APIs
* Vim
