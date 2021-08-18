# Issue (Beta)
Sometimes, if you're feverishly working on a project in your code editor or CLI, it can be **frustrating as hell** to keep switching to the browser to view task requirements in your issue tracker.

Plus, annoyingly, a lot of issue trackers do not have good support for keyboard shortcuts. You have to use the mouse a lot.

To combat all this, **Issue** is a simple CLI utility for viewing, editing and creating GitLab project issues _(Support for Jira, GitHub and more coming soon!)_.

# Installation
* Download the executable corresponding to your OS from the table below.

   | OS & Architecture | Link |
   |---|---|
   |Darwin (Mac OS) AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/darwin_amd64/issue) |
   |Darwin (Mac OS) ARM64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/darwin_arm64/Issue) |
   |Linux 386| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_386/Issue) |
   |Linux AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_amd64/issue) |
   |Linux ARM| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_arm/Issue) |
   |Linux ARM64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/linux_arm64/Issue) |
   |Windows 386| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/windows_386/Issue.exe) |
   |Windows AMD64| [Download](https://github.com/ArjArav98/Issue/raw/master/dist/windows_amd64/issue.exe) |

* Add the path of the directory, in which the executable is present, to your `PATH` environment variable.
* Run `issue version` to verify the successful installation!

# Usage

`issue list` - Lists all issues.

`issue list --assignee_username darth.vader69 --labels "Doing,Backend"` - Lists all issues after applying search parameters. For full list of search parameters and examples, click here.

`issue list --my-open-issues` - Shortcut which displays all issues assigned to you, which are open.

`issue list --my-issues` - Shortcut which displays all issues assigned to you.

------------

`issue show <issue_id>` - Displays the selected issue in detail.

`issue show --with-comments <issue_id>` - Displays the selected issue in detail, along with its comments.

`issue show --no-comments <issue_id>` - Displays the selected issue's comments only.

------------

`issue init` - Creates an empty config file in the current directory.

`issue version` - Displays the current version of the tool.
