package dao

import "vodstv/core/models"

//MatchDao ...
type MatchDao struct{}

//Get ...
func (d MatchDao) Get(id uint) (match models.Match, err error) {
	err = GetDB().First(&match, id).Related(&match.Vods).Error
	return match, err
}

//Find ...
func (d MatchDao) Find(q map[string]interface{}) (matchs []models.Match, err error) {
	err = GetDB().Where(q).Preload("Vods").Find(&matchs).Error
	return matchs, err
}

//Save ...
func (d MatchDao) Save(match *models.Match) (err error) {
	if GetDB().NewRecord(match) {
		err = GetDB().Create(&match).Error
	} else {
		err = GetDB().Save(&match).Updates(getUpdates(match)).Error
	}

	return err
}

//Delete ...
func (d MatchDao) Delete(id uint) error {
	match, err := d.Get(id)
	if err == nil {
		err = GetDB().Delete(&match).Error
	}
	return err
}
