package models

import "esvodsApi/forms"

//Article ...
type Article struct {
	ID        int64   `json:"id"`
	Watcher   Watcher `json:"-"`
	WatcherID int
	Title     string `json:"title"`
	Content   string `json:"content"`
	UpdatedAt int64  `json:"updated_at"`
	CreatedAt int64  `json:"created_at"`
}

//ArticleModel ...
type ArticleModel struct{}

//Create ...
func (m ArticleModel) Create(watcherID uint, form forms.ArticleForm) (articleID int64, err error) {
	// getDb := db.GetDB()

	// watcherModel := new(WatcherModel)

	// checkWatcher, err := watcherModel.One(watcherID)

	// if err != nil && checkWatcher.ID > 0 {
	// 	return 0, errors.New("Watcher doesn't exist")
	// }

	// _, err = getDb.Exec("INSERT INTO article(watcher_id, title, content, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", watcherID, form.Title, form.Content, time.Now().Unix(), time.Now().Unix())

	// if err != nil {
	// 	return 0, err
	// }

	// articleID, err = getDb.SelectInt("SELECT id FROM article WHERE watcher_id=$1 ORDER BY id DESC LIMIT 1", watcherID)

	return articleID, err
}

//One ...
func (m ArticleModel) One(watcherID uint, id int64) (article Article, err error) {
	// err = db.GetDB().SelectOne(&article, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS watcher FROM article a LEFT JOIN public.watcher u ON a.watcher_id = u.id WHERE a.watcher_id=$1 AND a.id=$2 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email LIMIT 1", watcherID, id)
	return article, err
}

//All ...
func (m ArticleModel) All(watcherID uint) (articles []Article, err error) {
	// _, err = db.GetDB().Select(&articles, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS watcher FROM article a LEFT JOIN public.watcher u ON a.watcher_id = u.id WHERE a.watcher_id=$1 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email ORDER BY a.id DESC", watcherID)
	return articles, err
}

//Update ...
func (m ArticleModel) Update(watcherID uint, id int64, form forms.ArticleForm) (err error) {
	// _, err = m.One(watcherID, id)

	// if err != nil {
	// 	return errors.New("Article not found")
	// }

	// _, err = db.GetDB().Exec("UPDATE article SET title=$1, content=$2, updated_at=$3 WHERE id=$4", form.Title, form.Content, time.Now().Unix(), id)

	return err
}

//Delete ...
func (m ArticleModel) Delete(watcherID uint, id int64) (err error) {
	// _, err = m.One(watcherID, id)

	// if err != nil {
	// 	return errors.New("Article not found")
	// }

	// _, err = db.GetDB().Exec("DELETE FROM article WHERE id=$1", id)

	return err
}
