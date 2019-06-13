package models

import (
	"errors"
)

func checkName(name string) error {
	if len(name) < 4 {
		return errors.New("todo name is too short")
	} else if len(name) > 20 {
		return errors.New("todo name is too long")
	} else {
		return nil
	}
}

func checkDescription(description string) error {
	if len(description) < 4 {
		return errors.New("todo description is too short")
	} else if len(description) > 150 {
		return errors.New("todo description is too long")
	} else {
		return nil
	}
}

func CheckTodoContent(todo Todo) error {
	err := checkName(todo.Name)
	if err != nil {
		return err
	}
	err = checkDescription(todo.Description)
	if err != nil {
		return err
	}
	return err
}
