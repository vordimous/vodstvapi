package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"esvodsApi/models"
	"log"
	"reflect"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// DbMigration ...
func DbMigration() {
	// getDb.AutoMigrate(&Vod{})
	GetDB().AutoMigrate(&models.Watcher{})
}

//WatcherSessionInfo ...
type WatcherSessionInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//GetWatcherID ...
func GetWatcherID(c *gin.Context) uint {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		return convertToUInt(watcherID)
	}
	return 0
}

//GetSessionWatcherInfo ...
func GetSessionWatcherInfo(c *gin.Context) (watcherSessionInfo WatcherSessionInfo) {
	session := sessions.Default(c)
	watcherID := session.Get("watcher_id")
	if watcherID != nil {
		watcherSessionInfo.ID = convertToUInt(watcherID)
		watcherSessionInfo.Name = session.Get("watcher_name").(string)
		watcherSessionInfo.Email = session.Get("watcher_email").(string)
	}
	return watcherSessionInfo
}

//JSONRaw ...
type JSONRaw json.RawMessage

//Value ...
func (j JSONRaw) Value() (driver.Value, error) {
	byteArr := []byte(j)
	return driver.Value(byteArr), nil
}

//Scan ...
func (j *JSONRaw) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []bytes"))
	}
	err := json.Unmarshal(asBytes, &j)
	if err != nil {
		return error(errors.New("Scan could not unmarshal to []string"))
	}
	return nil
}

//MarshalJSON ...
func (j *JSONRaw) MarshalJSON() ([]byte, error) {
	return *j, nil
}

//UnmarshalJSON ...
func (j *JSONRaw) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("json.RawMessage: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

func convertToUInt(number interface{}) uint {
	if reflect.TypeOf(number).String() == "int" {
		return uint(number.(int))
	}
	return number.(uint)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
