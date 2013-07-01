package semaphoreci

import (
	"encoding/json"
	"fmt"
)

type Project struct {
	Id        int
	HashId    string `json:"hash_id"`
	Name      string
	Owner     string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	branches  []Branch
	Client    *Client
}

type Branch struct {
	BrancName        string `json:"branch_name"`
	BranchURL        string `json:"branch_url"`
	BranchStatusURL  string `json:"branch_status_url"`
	BranchHistoryURL string `json:"branch_history_url"`
	ProjectName      string `json:"project_name"`
	BuildURL         string `json:"build_url"`
	BuildInfoURL     string `json:"build_info_url"`
	BuildNumber      int    `json:"build_number"`
	Result           string `json:"result"`
	StartedAt        string `json:"started_at"`
	FinishedAt       string `json:"finished_at"`
}

type Commit struct {
	Id         string
	URL        string
	AuthorName string `json:"author_name"`
	AuthorMail string `json:"author_mail"`
	Message    string
	Timestamp  string
}

type BranchStatus struct {
	Branch
	Commit Commit
}

type BranchHistory struct {
	Branch
	Builds []Build
}

type Build struct {
	BranchInfo
	Commit Commit
}

type BuildInfo struct {
	ProjectName string `json:"project_name"`
	BrancName   string `json:"branch_name"`
	Number      int
	Result      string
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	StartedAt   string `json:"started_at"`
	FinishedAt  string `json:"finished_at"`
	HtmlURL     string `json:"html_url"`
	Commits     []Commit
}

type BranchInfo struct {
	BuildURL     string `json:"build_url"`
	BuildInfoURL string `json:"build_info_url"`
	BuildNumber  int    `json:"build_number"`
	Result       string
	StartedAt    string `json:"started_at"`
	FinishedAt   string `json:"finished_at"`
}

type BuildLog struct {
	Threads      []Thread
	BuildInfoURL string `json:"build_info_url"`
}

type Thread struct {
	Number   int
	Commands []Command
}

type Command struct {
	Name       string
	Result     string
	Output     string
	Duration   string
	StartTime  string `json:"start_time"`
	FinishTime string `json:"finish_time"`
}

func (c *Client) Projects() ([]Project, error) {
	data := []Project{}
	body, _ := c.GetRequest("projects")
	err := json.Unmarshal(body, &data)
	return data, err
}

func (c *Client) Project(hash_id string) *Project {
	return &Project{HashId: hash_id, Client: c}
}

func (p *Project) Branches() (interface{}, error) {
	var data interface{}
	url := fmt.Sprintf("projects/%v/branches", p.HashId)
	body, _ := p.Client.GetRequest(url)
	err := json.Unmarshal(body, &data)
	return data, err
}

func (p *Project) BranchStatus(branch_id int) (BranchStatus, error) {
	data := BranchStatus{}
	url := fmt.Sprintf("projects/%v/%v/status", p.HashId, branch_id)
	body, _ := p.Client.GetRequest(url)
	err := json.Unmarshal(body, &data)
	return data, err
}

func (p *Project) BranchHistory(branch_id int) (BranchHistory, error) {
	data := BranchHistory{}
	url := fmt.Sprintf("projects/%v/%v", p.HashId, branch_id)
	body, _ := p.Client.GetRequest(url)
	err := json.Unmarshal(body, &data)
	return data, err
}

func (p *Project) BuildInfo(branch_id, build_num int) (BuildInfo, error) {
	data := BuildInfo{}
	url := fmt.Sprintf("projects/%v/%v/builds/%v", p.HashId, branch_id, build_num)
	body, _ := p.Client.GetRequest(url)
	err := json.Unmarshal(body, &data)
	return data, err
}

func (p *Project) BuildLog(branch_id, build_num int) (BuildLog, error) {
	data := BuildLog{}
	url := fmt.Sprintf("projects/%v/%v/builds/%v/log", p.HashId, branch_id, build_num)
	body, _ := p.Client.GetRequest(url)
	err := json.Unmarshal(body, &data)
	return data, err
}
