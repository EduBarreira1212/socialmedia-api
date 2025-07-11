package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	NickName  string    `json:"nickName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is mandatory")
	}

	if user.NickName == "" {
		return errors.New("nickname is mandatory")
	}

	if user.Email == "" {
		return errors.New("e-mail is mandatory")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail invalid")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password is mandatory")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.NickName = strings.TrimSpace(user.NickName)
	user.Email = strings.TrimSpace(user.Email)
}
