package gstruct

import (
	"fmt"
	"testing"
)

type RequestCreateArticle struct {
	TagNames    []string `json:"tag_names" validate:"required,dive,max=25"`
	AuthorId    string   `json:"author_id" validate:"required,ulid"`
	Title       string   `json:"title" validate:"required,max=100,min=15"`
	Description string   `json:"description" validate:"required,max=255,min=25"`
	Body        string   `json:"body" validate:"required,min=50"`
}

func TestMarshalAndCencoredTag(t *testing.T) {
	req := &RequestCreateArticle{
		TagNames:    []string{"tag1", "tag2"},
		AuthorId:    "author123",
		Title:       "Sample Title",
		Description: "Sample Description",
		Body:        "Sample Body",
	}

	result, err := MarshalAndCencoredTag(req, "cencored")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
