package forms

import (
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//FeedForm ...
type FeedForm struct {
	ModelForm
	ID    uint `json:"id"`
	Name  string
	Type  string
	Regex string
}

//FeedSearch ...
type FeedSearch struct {
	Name string
	Type string
}

//ToModel ...
func (f FeedForm) ToModel(t *models.Feed) (err error) {
	m := structs.Map(f)
	t.FillStruct(m)
	return err
}
