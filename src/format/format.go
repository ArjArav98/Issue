package format

import (
	"fmt"
	"strings"
	"github.com/ArjArav98/Issue/src/types"
)

/*-------------------*/
/* EXPOSED FUNCTIONS */
/*-------------------*/

func BeautifyIssue (issue types.Issue) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf(`
- ISSUE NO. - %v
- TITLE     - %s
- STATE     - %s
- LABELS    - %s
- ASSIGNEE  - %s (%s)
- AUTHOR    - %s (%s) created this on %s.
- WEB URL   - %s

###############
# DESCRIPTION #
###############

%s
	`, issue.Iid, issue.Title, issue.State, strings.Join(issue.Labels, ", "),
	   issue.Assignee.Name, issue.Assignee.Username, issue.Author.Name, issue.Author.Username,
	   issue.CreatedAt, issue.WebUrl, issue.Description))

	return builder.String()
}

func BeautifyComments (comments []types.Comment) string {
	var builder strings.Builder

	builder.WriteString(`
############
# COMMENTS #
############
	`)

	for iter:=0; iter<len(comments); iter++ {
		builder.WriteString(BeautifyComment(comments[iter]))
	}

	return builder.String()
}

func BeautifyComment (comment types.Comment) string {
	var builder strings.Builder

	if !comment.SystemGenerated {
		builder.WriteString(fmt.Sprintf(`
%v COMMENTED AT %v,

"%v"

=+=+=+=+=+=+=
=+=+=+=+=+=+=
		`, strings.ToUpper(comment.Author.Name), comment.UpdatedAt, comment.Body))

	}

	return builder.String()
}
