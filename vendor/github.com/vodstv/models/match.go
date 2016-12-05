package models

import "time"

//Match ...
type Match struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	Title     string     `json:"title"`
	Vods      []Vod      `json:"vods"`
}

//BeforeSave ...
// func (m Match) BeforeSave() (err error) {
//     var vs = []Match
//     dao.GetDB().Model(&m).Related(&vs)
//     var tags = make(map[string]bool)
//     for _, v := range vs) {
// 		var ts = []Match
//         dao.GetDB().Model(&v).Related(&ts)
//         for _, t := range ts) {
//             tags[t.Name] = true
//         }
// 	}
// 	return
// }
