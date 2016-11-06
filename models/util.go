package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"esvodsApi/db"
	"reflect"
)

// DbMigration ...
func DbMigration() {
	getDb := db.GetDB()
	// getDb.AutoMigrate(&Article{})
	getDb.AutoMigrate(&Watcher{})
}

//WatcherSessionInfo ...
type WatcherSessionInfo struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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

//ConvertToUInt ...
func ConvertToUInt(number interface{}) uint {
	if reflect.TypeOf(number).String() == "int" {
		return uint(number.(int))
	}
	return number.(uint)
}
