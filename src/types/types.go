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
	Created_at string
	Updated_at string
	Author User
	Assignee User
	Labels []string
	WebUrl string
}

/*------------------------------*/
/* JSON UNMARSHALLING FUNCTIONS */
/*------------------------------*/

func (issue *Issue) FromJson (stringContent []byte) error {
	err := json.Unmarshal(stringContent, issue)
	if err != nil {
		return errors.New("the raw data could not be parsed into JSON")
	}

	return nil
}
