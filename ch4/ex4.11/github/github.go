// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github


const IssuesURL = "https://api.github.com/repos/samael65535/gopl_exercise/issues/"

type Issue struct {
	Id        uint64 `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	State     string `json:"state,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}

var USERNAME string
var PASSWORD string
//!-
