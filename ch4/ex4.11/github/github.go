// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github


const IssuesURL = "https://api.github.com/repos/samael65535/gopl_exercise/issues"

type Issue struct {
	Id        int `json:"id,omitempty"`
	Title     string `json:"title"`
	Body      string `json:body,omitempty`
}

var USERNAME string
var PASSWORD string
//!-
