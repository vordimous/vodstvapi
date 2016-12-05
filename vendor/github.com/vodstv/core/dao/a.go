package dao

import (
	"strings"

	"github.com/fatih/structs"
)

func getQuery(s *structs.Struct) map[string]interface{} {
	q := make(map[string]interface{})
	for _, f := range s.Fields() {
		v := f.Value()
		if v != nil && v != "" {
			q[strings.ToLower(f.Name())] = v
		}
	}
	return q
}

func getUpdates(s interface{}) map[string]interface{} {
	m := structs.Map(s)
	delete(m, "ID")
	delete(m, "CreatedAt")
	delete(m, "UpdatedAt")
	delete(m, "DeletedAt")
	return m
}
