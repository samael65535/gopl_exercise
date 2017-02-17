// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

package github

import (
	"encoding/json"
	"bytes"
	"net/http"
	"fmt"
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
	fmt.Println(string(data))
	client := &http.Client{}

	req, errNet := http.NewRequest(http.MethodPost, IssuesURL, bytes.NewReader(data))
	defer req.Body.Close()
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

func LoadIssue(id uint64) {

}

func CloseIssue(id string) (error){
	issue := Issue{
		State: "closed",
	}
	data, err :=  json.Marshal(issue)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	client := &http.Client{}
	fmt.Println(IssuesURL + id)
	req, errNet := http.NewRequest(http.MethodPost, IssuesURL + id, bytes.NewReader(data))
	defer req.Body.Close()
	if errNet != nil {
		return errNet
	}
	req.Header.Set("Content-Type", "application/json");
	req.SetBasicAuth(USERNAME, PASSWORD)
	_, err = client.Do(req)

	if errNet != nil {
		return errNet
	}
	fmt.Println(req.Body.Text())
	return nil
}

func EditIssue(title string, body string, id uint64) {

}


//!-
