package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const base = "https://api.github.com/repos/"

type RepositoryInfo struct {
	RepoName      string
	Issues        []Issue
	Milestones    []Milestone
	Collaborators []Collaborator
}

type Issue struct {
	Id        int       `json:"id"`
	URL       string    `json:"html_url"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

type Milestone struct {
	Id          int       `json:"id"`
	URL         string    `json:"html_url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
}

type Collaborator struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"html_url"`
}

func GetRepositoryInfo(token, user, repo string) (*RepositoryInfo, error) {
	issues, err := ListIssue(user, repo)
	if err != nil {
		return nil, err
	}
	milestones, err := ListMilestone(user, repo)
	if err != nil {
		return nil, err
	}
	collaborators, err := ListCollaborator(token, user, repo)
	if err != nil {
		return nil, err
	}
	return &RepositoryInfo{
		RepoName:      repo,
		Issues:        issues,
		Milestones:    milestones,
		Collaborators: collaborators,
	}, nil
}

func ListIssue(user, repo string) ([]Issue, error) {
	url := base + user + "/" + repo + "/issues"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func ListMilestone(user, repo string) ([]Milestone, error) {
	url := base + user + "/" + repo + "/milestones"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// Collaboratorの取得
func ListCollaborator(token, user, repo string) ([]Collaborator, error) {
	url := base + user + "/" + repo + "/collaborators"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []Collaborator
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
