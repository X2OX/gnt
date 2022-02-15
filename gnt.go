package gnt

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/webhooks/v6/github"
)

var hook *github.Webhook

func init() {
	var err error
	if hook, err = github.New(github.Options.Secret(os.Getenv("GNT_GITHUB_SECRET"))); err != nil {
		panic(err)
	}
}

func Parse(r *http.Request) (string, string, error) {
	if r == nil {
		return "", "", fmt.Errorf("nil request")
	}

	payload, err := hook.Parse(r, github.CommitCommentEvent,
		github.IssueCommentEvent, github.IssuesEvent,
		github.LabelEvent, github.MilestoneEvent, github.MetaEvent,
		github.ProjectCardEvent, github.ProjectColumnEvent, github.ProjectEvent,
		github.PullRequestEvent, github.PullRequestReviewEvent, github.PullRequestReviewCommentEvent,
		github.PushEvent, github.ReleaseEvent, github.RepositoryEvent)
	if err != nil {
		return "", "", err
	}
	switch payload := payload.(type) {
	case github.IssueCommentPayload:
		return fmt.Sprintf(`\#%d *%s*%s\#IssueComment \#%s \#%s`, payload.Issue.Number, EscapedMarkdownV2(payload.Issue.Title),
			FilterBody(payload.Comment.Body), payload.Action, payload.Sender.Login), payload.Issue.HTMLURL, nil
	case github.IssuesPayload:
		return fmt.Sprintf(`\#%d *%s*%s\#Issues \#%s \#%s`, payload.Issue.Number, EscapedMarkdownV2(payload.Issue.Title),
			FilterBody(payload.Issue.Body), payload.Action, payload.Sender.Login), payload.Issue.HTMLURL, nil
	}
	return FilterBody(fmt.Sprintf("%T: not processed", payload)), "", nil
}
