package shield

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

/* Extract the mapping between JSON names and field name of a structure
 */
func extractTags(s interface{}) map[string]string {
	m := make(map[string]string)
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		m[field.Tag.Get("json")] = field.Name
	}
	return m
}

func ValidateFields(js []byte, s interface{}, ga []string) error {
	t := extractTags(s)
	st := reflect.TypeOf(s)
	m := make(map[string]string)
	json.Unmarshal(js, &m)
	var authorized bool
	for k, _ := range m {
		// Field is not present in object
		if _, ok := t[k]; ok == false {
			return errors.New("Field " + k + " not supported")
		}
		// Call authorized to operate on field
		f, _ := st.FieldByName(t[k])
		authorized = false
		for _, va := range ga {
			for _, vt := range strings.Split(f.Tag.Get("groups"), ",") {
				if va == vt {
					authorized = true
				}
			}
		}
		if authorized == false {
			return errors.New("Field " + k + " not writable")
		}
	}
	return nil
}
