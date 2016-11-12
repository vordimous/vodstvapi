package forms

import (
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//TagForm ...
type TagForm struct {
	ModelForm
	ID    uint `json:"id"`
	Name  string
	Type  string
	Regex string
}

//TagSearch ...
type TagSearch struct {
	Name string
	Type string
}

//ToModel ...
func (f TagForm) ToModel(t *models.Tag) (err error) {
	m := structs.Map(f)
	t.FillStruct(m)
	return err
}
