package forms

import (
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//VodForm ...
type VodForm struct {
	ModelForm
	ID      uint
	Title   string
	Content string
}

//VodSearch ...
type VodSearch struct {
	Title   string
	Content string
}

//ToModel ...
func (f VodForm) ToModel(v *models.Vod) (err error) {
	m := structs.Map(f)
	v.FillStruct(m)
	return err
}
