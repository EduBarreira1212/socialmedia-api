package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.Validate(); err != nil {
		return err
	}

	post.Format()
	return nil
}

func (post *Post) Validate() error {
	if post.Title == "" {
		return errors.New("title is required")
	}

	if post.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (post *Post) Format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
