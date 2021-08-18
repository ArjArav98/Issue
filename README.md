# Issue (Beta)
Sometimes, if you're feverishly working on a project in your code editor or CLI, it can be **frustrating as hell** to keep switching to the browser to view task requirements in your issue tracker.

Plus, annoyingly, a lot of issue trackers do not have good support for keyboard shortcuts. You have to use the mouse a lot.

To combat all this, **Issue** is a simple CLI utility for viewing, editing and creating GitLab project issues _(Support for Jira, GitHub and more coming soon!)_.

# Installation
* Clone this repository.
* In the root directory, run `go build -o issue main.go`.
    * Add the current directory path to your `PATH` environment variable.
* Run `issue init` to verify the successful installation!

# Usage
* **Listing all issues** - `issue list`
   * For filtering this list, add search parameters. Examples below.
      * `issue list --assignee_username darthvader --labels "Doing,Backend"`
      * `issue list --state opened --created_before 2019-03-15T08:00:00Z`
      * You can add as many parameters as you want. For the full list, click here.
   * There are two available shortcuts to avoid typing long searches.
      * 
------------
* **Seeing a single issue in detail** - `issue show <issue_id>`
   * Seeing issue with comments/notes - `issue show --with-comments <issue_id>`
   * Seeing only comments - `issue show --no-comments <issue_id>`
   * 
