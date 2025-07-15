package iamtk

import (
	"reflect"

	"github.com/fale/sheriff"
	"github.com/google/uuid"
)

func marshalItem(ps *[]Permission, item interface{}, desiredCapabilities []string, field string) (map[string]interface{}, error) {
	// Get the value of the field
	var id uuid.UUID
	if len(field) > 0 {
		fieldValue := reflect.ValueOf(item).FieldByName(field)
		if fieldValue.Kind() == reflect.String {
			t := fieldValue.Interface().(string)
			id, _ = uuid.Parse(t)
		} else {
			id = fieldValue.Interface().(uuid.UUID)
		}
	} else {
		id = WildCard
	}
	permissions := PertinentCapabilities(ps, id, desiredCapabilities)

	// If no permissions, nothing should be returned
	if len(permissions) == 0 {
		return map[string]interface{}{}, nil
	}

	// Clean with sheriff
	o := &sheriff.Options{
		Groups: []sheriff.Group{
			{Values: permissions},
		},
	}

	m, err := sheriff.Marshal(o, item)
	if err != nil {
		return nil, err
	}
	out := m.(map[string]interface{})
	return out, nil
}

func Marshal(ps *[]Permission, items interface{}, desiredCapabilities []string, field string, limit int) (interface{}, error) {
	val := reflect.ValueOf(items)
	if val.Kind() == reflect.Slice {
		var out []map[string]interface{}
		var o map[string]interface{}
		var err error
		for i := 0; i < val.Len(); i++ {
			item := val.Index(i).Interface()
			o, err = marshalItem(ps, item, desiredCapabilities, field)
			if err != nil {
				return nil, err
			}
			if len(out) >= limit {
				break
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
	return marshalItem(ps, items, desiredCapabilities, field)
}
