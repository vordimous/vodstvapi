package dao

import "vodstv/core/models"

//VodDao ...
type VodDao struct{}

//Get ...
func (d VodDao) Get(id uint) (vod models.Vod, err error) {
	err = GetDB().First(&vod, id).Related(&vod.Tags, "Tags").Error
	return vod, err
}

//Find ...
func (d VodDao) Find(q map[string]interface{}) (vods []models.Vod, err error) {
	err = GetDB().Where(q).Preload("Tags").Find(&vods).Error
	return vods, err
}

//Query ...
func (d VodDao) Query(tags []uint) (vods []models.Vod, err error) {
	vodIDs := []uint{}
	dbQuery := db.Table("vod_tags")
	if tags != nil && len(tags) > 0 {
		dbQuery = dbQuery.Where("tag_id in (?)", tags)
	}
	err = dbQuery.Pluck("DISTINCT(vod_id)", &vodIDs).Error
	if err == nil {
		err = GetDB().Where("id in (?)", vodIDs).Preload("Tags").Find(&vods).Error
	}
	return vods, err
}

//Save ...
func (d VodDao) Save(vod *models.Vod) (err error) {
	if GetDB().NewRecord(vod) {
		err = GetDB().Create(&vod).Error
	} else {
		err = GetDB().Save(&vod).Updates(getUpdates(vod)).Error
	}

	return err
}

//Delete ...
func (d VodDao) Delete(id uint) (vod models.Vod, err error) {
	vod, err = d.Get(id)
	if err == nil {
		err = GetDB().Delete(&vod).Error
	}
	return vod, err
}
