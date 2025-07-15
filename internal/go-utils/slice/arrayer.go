package slice

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Arrayer(items []interface{}, columns []string) [][]string {
	var out [][]string

	for _, s := range items {
		var o []string
		for _, column := range columns {
			o = append(o, findColumn(s, column))
		}
		out = append(out, o)
	}

	return out
}

func findColumn(s interface{}, column string) string {
	var remainer string
	if strings.Contains(column, "/") {
		s := strings.Split(column, "/")
		remainer = strings.Join(s[1:], "/")
		column = s[0]
	}
	var st reflect.Type
	var sv reflect.Value
	if reflect.ValueOf(s).Kind() == reflect.Ptr {
		st = reflect.TypeOf(s).Elem()
		sv = reflect.ValueOf(s).Elem()
	} else {
		st = reflect.TypeOf(s)
		sv = reflect.ValueOf(s)
	}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			if json == column {
				switch field.Type.String() {
				case "bool":
					return strconv.FormatBool(sv.FieldByName(field.Name).Interface().(bool))
				case "int":
					return strconv.Itoa(sv.FieldByName(field.Name).Interface().(int))
				case "int64":
					return strconv.FormatInt(sv.FieldByName(field.Name).Interface().(int64), 10)
				case "time.Time":
					return sv.FieldByName(field.Name).Interface().(time.Time).Format("2006-01-02 15:04:05 UTC")
				case "uint64":
					return strconv.FormatUint(sv.FieldByName(field.Name).Interface().(uint64), 10)
				case "uuid.UUID":
					return sv.FieldByName(field.Name).Interface().(uuid.UUID).String()
				default:
					if len(remainer) > 0 {
						return findColumn(sv.FieldByName(field.Name).Interface(), remainer)
					}
					return sv.FieldByName(field.Name).String()
				}
			}
		} else {
			log.Println(field)
		}
	}
	return ""
}
