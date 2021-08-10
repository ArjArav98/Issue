package types

import (
	"errors"
	"encoding/json"
)

/*-------*/
/* TYPES */
/*-------*/

type User struct {
	Id uint64
	Name string
	Username string
}

type Issue struct {
	Id uint64
	Iid uint64
	Title string
	Description string
	State string
	Created_At string
	Updated_At string
	Author User
	Assignee User
	Labels []string
	Web_Url string
	Project_Id uint64
}

type Comment struct {
	Id uint64
	Body string
	Author User
	System bool
	Created_At string
	Updated_At string
}

type Project struct {
	Id uint64
	Name string
	Description string
	Visibility string
	Web_Url string
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
