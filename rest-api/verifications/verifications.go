package verifications

import (
	"errors"

	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
)

func checkName(todo models.Todo) error {
	if len(todo.Name) < 4 {
		return errors.New("todo name is too short")
	} else if len(todo.Name) > 20 {
		return errors.New("todo name is too long")
	} else {
		return nil
	}
}

func checkDescription(todo models.Todo) error {
	if len(todo.Description) < 4 {
		return errors.New("todo description is too short")
	} else if len(todo.Description) > 150 {
		return errors.New("todo description is too long")
	} else {
		return nil
	}
}

func CheckTodoContent(todo models.Todo) error {
	err := checkName(todo)
	if err != nil {
		return err
	}
	err = checkDescription(todo)
	if err != nil {
		return err
	}
	return err
}
