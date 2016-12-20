package sess

import (
	"strconv"

	"vodstv/core"
	"vodstv/core/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//WatcherSessionInfo ...
type WatcherSessionInfo struct {
	ID      uint          `json:"id"`
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	IsAdmin bool          `json:"isAdmin"`
	Feeds   []models.Feed `json:"feeds"`
}

//GetWatcherID ...
func GetWatcherID(c *gin.Context) uint {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		return core.ConvertToUInt(watcherID)
	}
	return 0
}

//GetSessionWatcherInfo ...
func GetSessionWatcherInfo(c *gin.Context) (watcherSessionInfo WatcherSessionInfo) {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		watcherSessionInfo.ID = core.ConvertToUInt(watcherID)
		watcherSessionInfo.Name = session.Get("watcher_name").(string)
		watcherSessionInfo.Email = session.Get("watcher_email").(string)

		if session.Get("watcher_is_admin") != nil {
			isAdmin, _ := strconv.ParseBool(session.Get("watcher_is_admin").(string))
			watcherSessionInfo.IsAdmin = isAdmin
		} else {
			watcherSessionInfo.IsAdmin = true
		}
	}
	return watcherSessionInfo
}
