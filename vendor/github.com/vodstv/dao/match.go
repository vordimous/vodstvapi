package dao

import "github.com/vodstv/core/models"

//MatchDao ...
type MatchDao struct{}

//Get ...
func (d MatchDao) Get(id uint) (match models.Match, err error) {
	err = GetDB().First(&match, id).Related(&match.Vods).Error
	return match, err
}

//Find ...
func (d MatchDao) Find(q map[string]interface{}) (matchs []models.Match, err error) {
	err = db.Where(q).Find(&matchs).Error
	return matchs, err
}

//Save ...
func (d MatchDao) Save(match *models.Match) (err error) {
	if GetDB().NewRecord(match) {
		err = GetDB().Create(&match).Error
	} else {
		err = GetDB().Save(&match).Error
	}

	return err
}

//Delete ...
func (d MatchDao) Delete(id uint) error {
	match, err := d.Get(id)
	if err == nil {
		err = db.Delete(&match).Error
	}
	return err
}
