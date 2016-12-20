package dao

import "vodstv/core/models"

//FeedDao ...
type FeedDao struct{}

//Get ...
func (d FeedDao) Get(id uint) (feed models.Feed, err error) {
	err = GetDB().First(&feed, id).Related(&feed.Tags, "Tags").Error
	return feed, err
}

//Find ...
func (d FeedDao) Find(q map[string]interface{}) (feeds []models.Feed, err error) {
	err = GetDB().Where(q).Preload("Tags").Find(&feeds).Error
	return feeds, err
}

//Save ...
func (d FeedDao) Save(feed *models.Feed) (err error) {
	if GetDB().NewRecord(feed) {
		err = GetDB().Create(&feed).Error
	} else {
		err = GetDB().Save(&feed).Updates(getUpdates(feed)).Error
	}

	return err
}

//Delete ...
func (d FeedDao) Delete(id uint) error {
	feed, err := d.Get(id)
	if err == nil {
		err = GetDB().Delete(&feed).Error
	}
	return err
}
