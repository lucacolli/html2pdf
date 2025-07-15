package shield

import (
	"reflect"

	liip "github.com/liip/sheriff"
)

func marshalItem(realm string, authorizations Authorizations, mode string, item interface{}, field string) (map[string]interface{}, error) {
	// Get the value of the field
	fieldValue := reflect.ValueOf(item).FieldByName(field).String()
	permissions := GetPermissions(realm, authorizations, mode, fieldValue)

	// Clean with sheriff
	o := &liip.Options{
		Groups: permissions,
	}

	m, err := liip.Marshal(o, item)
	if err != nil {
		return nil, err
	}

	out := m.(map[string]interface{})
	return out, nil
}

func Marshal(realm string, authorizations Authorizations, mode string, items interface{}, field string) (interface{}, error) {
	val := reflect.ValueOf(items)
	if val.Kind() == reflect.Slice {
		var out []map[string]interface{}
		var o map[string]interface{}
		var err error
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			o, err = marshalItem(realm, authorizations, mode, item, field)
			if err != nil {
				return nil, err
			}
			if len(o) != 0 {
				out = append(out, o)
			}
		}
		if len(out) == 0 {
			return []map[string]string{}, nil
		}
		return out, nil
	}
	return marshalItem(realm, authorizations, mode, items, field)

}
