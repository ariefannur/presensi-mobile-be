package models

import (
	"net/url"
	"regexp"
)

type User struct {
	Id       string `db:"id" json:"id" validate:"required"`
	Name     string `db:"title" json:"title" validate:"required,lte=100"`
	Email    string `db:"email" json:"email" validate:"required,lte=100"`
	Password string `db:"password" json:"password" validate:"required,lte=100"`
	UserType string `db:"type" json:"type" validate:"required,lte=100"`
}

func (u User) IsValid() (errs url.Values) {
	if u.Name == "" {
		errs.Add("name", "The name is required!")
	}

	if u.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if !emailRegex.MatchString(u.Email) {
		errs.Add("email", "The email field should be a valid email address!")
	}

	if len(errs) == 0 {
		return nil
	} else {
		return errs
	}
}

type UserType struct {
	Name string `json:"name"`
}
