package http

import (
	"encoding/json"
	"errors"
	"time"
)

/*
{
   "complete": true,
}
*/

type CompleteTaskDTO struct {
	Complete bool
}

// DTO == data transfer object
type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
	E       *ErrorDTO
}

func (e ErrorDTO) toString() string {
	newE := ErrorDTO{}
	newE.E = &newE

	// MarshalIndent делает красивые отступы
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	
	return string(b)
}