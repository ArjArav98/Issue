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
	   issue.Created_At, issue.Web_Url, issue.Description))

	return builder.String()
}
