package types

import (
	"errors"
	"encoding/json"
)

/*-------*/
/* TYPES */
/*-------*/

type User struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
}

type Issue struct {
	Id uint64 `json:"id"`
	Iid uint64 `json:"iid"`
	Title string `json:"title"`
	Description string `json:"description"`
	State string `json:"state"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Author User `json:"author"`
	Assignee User `json:"assignee"`
	Labels []string `json:"labels"`
	WebUrl string `json:"web_url"`
	ProjectId uint64 `json:"project_id"`
}

type Comment struct {
	Id uint64 `json:"id"`
	Body string `json:"body"`
	Author User `json:"author"`
	System bool `json:"system"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Project struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Visibility string `json:"visibility"`
	WebUrl string `json:"web_url"`
}

/*------------------------------*/
/* JSON UNMARSHALLING FUNCTIONS */
/*------------------------------*/

func (issue *Issue) FromJson (stringContent []byte) error {
	err := json.Unmarshal(stringContent, issue)
	if err != nil {
		return errors.New("the raw data for an Issue could not be parsed into JSON")
	}

	return nil
}

func (comment *Comment) FromJson (stringContent []byte) error {
	err := json.Unmarshal(stringContent, comment)
	if err != nil {
		return errors.New("the raw data for a Comment could not be parsed into JSON")
	}

	return nil
}

func (project *Project) FromJson (stringContent []byte) error {
	err := json.Unmarshal(stringContent, project)
	if err != nil {
		return errors.New("the raw data for a Project could not be parsed into JSON")
	}

	return nil
}

/* We require a different func signature since []Comment isn't a defined type. */
func CommentsFromJson (stringContent []byte, comments *[]Comment) error {
	err := json.Unmarshal(stringContent, comments)
	if err != nil {
		return errors.New("the raw data for Comments could not be parsed into JSON")
	}

	return nil
}
