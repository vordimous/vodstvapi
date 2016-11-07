package dao

import (
	"esvodsApi/forms"
	"esvodsApi/models"
)

//VodDao ...
type VodDao struct{}

//Create ...
func (d VodDao) Create(form forms.VodForm) (vodID uint, err error) {
	// getDb := db.GetDB()

	// watcherModel := new(WatcherModel)

	// checkWatcher, err := watcherModel.One(watcherID)

	// if err != nil && checkWatcher.ID > 0 {
	// 	return 0, errors.New("Watcher doesn't exist")
	// }

	// _, err = getDb.Exec("INSERT INTO vod(watcher_id, title, content, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", watcherID, form.Title, form.Content, time.Now().Unix(), time.Now().Unix())

	// if err != nil {
	// 	return 0, err
	// }

	// vodID, err = getDb.SelectInt("SELECT id FROM vod WHERE watcher_id=$1 ORDER BY id DESC LIMIT 1", watcherID)

	return vodID, err
}

//One ...
func (d VodDao) One(id uint) (vod models.Vod, err error) {
	// err = db.GetDB().SelectOne(&vod, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS watcher FROM vod a LEFT JOIN public.watcher u ON a.watcher_id = u.id WHERE a.watcher_id=$1 AND a.id=$2 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email LIMIT 1", watcherID, id)
	return vod, err
}

//All ...
func (d VodDao) All() (vods []models.Vod, err error) {
	// _, err = db.GetDB().Select(&vods, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS watcher FROM vod a LEFT JOIN public.watcher u ON a.watcher_id = u.id WHERE a.watcher_id=$1 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email ORDER BY a.id DESC", watcherID)
	return vods, err
}

//Update ...
func (d VodDao) Update(form forms.VodForm) (err error) {
	// _, err = d.One(watcherID, id)

	// if err != nil {
	// 	return errors.New("models.Vod not found")
	// }

	// _, err = db.GetDB().Exec("UPDATE vod SET title=$1, content=$2, updated_at=$3 WHERE id=$4", form.Title, form.Content, time.Now().Unix(), id)

	return err
}

//Delete ...
func (d VodDao) Delete(id uint) (err error) {
	// _, err = d.One(watcherID, id)

	// if err != nil {
	// 	return errors.New("models.Vod not found")
	// }

	// _, err = db.GetDB().Exec("DELETE FROM vod WHERE id=$1", id)

	return err
}
