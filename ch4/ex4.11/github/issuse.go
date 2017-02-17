// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"bytes"
	"net/http"
)

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) {
	// q := url.QueryEscape(strings.Join(terms, " "))

	// resp, err := http.Get(IssuesURL + "?q=" + q)
	// if err != nil {
	//	return nil, err
	// }
	// if resp.StatusCode != http.StatusOK {
	//	resp.Body.Close()
	//	return nil, fmt.Errorf("search query failed: %s", resp.Status)
	// }

	// var result IssuesSearchResult
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	//	resp.Body.Close()
	//	return nil, err
	// }
	// resp.Body.Close()
	// return &result, nil
}


func CreateIssue(issue *Issue) (error){
	data, err :=  json.Marshal(*issue)
	if err != nil {
		return err
	}

	client := &http.Client{}

	req, errNet := http.NewRequest(http.MethodPost, IssuesURL, bytes.NewReader(data))
	if errNet != nil {
		return errNet
	}
	req.Header.Set("Content-Type", "application/json");
	req.SetBasicAuth(USERNAME, PASSWORD)
	_, err = client.Do(req)

	if errNet != nil {
		return errNet
	}
	return nil
}

func LoadIssues(id uint64) {

}

func CloseIssues(id uint64) (bool){
	return false
}

func EditIssues(title string, body string, id uint64) {

}


//!-
