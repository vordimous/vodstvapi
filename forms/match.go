package forms

import (
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//MatchForm ...
type MatchForm struct {
	ModelForm
	ID    uint `json:"id"`
	Title string
}

//MatchVodAsc ...
type MatchVodAsc struct {
	MatchID uint `json:"matchId"`
	VodID   uint `json:"vodId"`
}

//MatchSearch ...
type MatchSearch struct {
	Name string
}

//ToModel ...
func (f MatchForm) ToModel(t *models.Match) (err error) {
	m := structs.Map(f)
	t.FillStruct(m)
	return err
}
