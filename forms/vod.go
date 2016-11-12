package forms

import (
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//VodForm ...
type VodForm struct {
	ModelForm `json:"-"`
	ID        uint
	Title     string
	Content   string
	VideoKey  string
	VideoURL  string
	VideoSrc  string
	VideoDate string
	ThumbURL  string
}

//VodTagAsc ...
type VodTagAsc struct {
	VodID uint `json:"vodId"`
	TagID uint `json:"tagId"`
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
