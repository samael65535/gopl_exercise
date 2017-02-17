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

func LoadIssue(id string) (*Issue, error){
	resp, err := http.Get(IssuesURL + id)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
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
	return nil
}

func EditIssue(title string, body string, id string) (error){
	issue := Issue{
		Title: title,
		Body: body,
	}
	data, err :=  json.Marshal(issue)
	if err != nil {
		return err
	}
	client := &http.Client{}
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
	return nil
}


//!-
