package core

import (
	"log"
	"reflect"
)

//ConvertToUInt ...
func ConvertToUInt(number interface{}) uint {
	if reflect.TypeOf(number).String() == "int" {
		return uint(number.(int))
	}
	return number.(uint)
}

//CheckErr ...
func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
