package models

// Repository - Repository info
type Repository struct {
	Name    string        `json:"name"`
	Avatar  string        `json:"avatar"`
	URL     string        `json:"url"`
	Commits []interface{} `json:"commits"`
}

// Commit - Commit info
type Commit struct {
	ID              string        `json:"id"`
	DisplayID       string        `json:"displayId"`
	AuthorTimestamp string        `json:"authorTimestamp"`
	Merge           bool          `json:"merge"`
	Files           []interface{} `json:"files"`
}

// Branch - Branch info
type Branch struct {
	Name                 string     `json:"name"`
	URL                  string     `json:"url"`
	CreatePullRequestURL string     `json:"createPullRequestUrl"`
	Repository           Repository `json:"repository"`
	LastCommit           Commit     `json:"lastCommit"`
}

// User - User info
type User struct {
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Approved bool   `json:"approved"`
}

// Source - Source info
type Source struct {
	Branch string `json:"branch"`
	URL    string `json:"url"`
}

// PullRequest - Pull Requests Info
type PullRequest struct {
	Author       User   `json:"author"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	CommentCount int    `json:"commentCount"`
	Source       Source `json:"source"`
	Destination  Source `json:"destination"`
	Reviewers    []User `json:"reviewers"`
	Status       string `json:"status"`
	URL          string `json:"url"`
	LastUpdate   string `json:"lastUpdate"`
}

// Instance - Instance Info
type Instance struct {
	SingleInstance bool   `json:"singleInstance"`
	Name           string `json:"name"`
	TypeName       string `json:"typeName"`
	ID             string `json:"id"`
	Type           string `json:"type"`
	BaseURL        string `json:"baseUrl"`
}

// Detail - Detail Info
type Detail struct {
	Branches     []Branch      `json:"branches"`
	PullRequests []PullRequest `json:"pullRequests"`
	Repositories []interface{} `json:"repositories"`
	Instance     Instance      `json:"_instance"`
}

// DevelopStatus - Develop Status info
type DevelopStatus struct {
	Errors []interface{} `json:"errors"`
	Detail []Detail      `json:"detail"`
}
