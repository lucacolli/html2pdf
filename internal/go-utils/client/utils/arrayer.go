package utils

import (
	"reflect"
	"time"
)

func Arrayer(items []interface{}, columns []string) [][]string {

	var out [][]string

	for _, s := range items {
		var o []string
		for _, column := range columns {
			st := reflect.TypeOf(s)
			for i := 0; i < st.NumField(); i++ {
				field := st.Field(i)
				if json, ok := field.Tag.Lookup("json"); ok {
					if json == column {
						if field.Type.String() == "time.Time" {
							o = append(o, reflect.ValueOf(s).FieldByName(field.Name).Interface().(time.Time).Format("2006-01-02 15:04:05 UTC"))
						} else {
							o = append(o, reflect.ValueOf(s).FieldByName(field.Name).String())
						}
					}
				}
			}
		}
		out = append(out, o)
	}

	return out
}
