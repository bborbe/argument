// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// ParseEnv parse the os.Environ() to data struct.
func ParseEnv(data interface{}, environ []string) error {
	values, err := envToValues(data, environ)
	if err != nil {
		return err
	}
	return fill(data, values)
}

func envToValues(data interface{}, environ []string) (map[string]interface{}, error) {
	var err error
	envValues := make(map[string]string)
	for _, env := range environ {
		for i := 0; i < len(env); i++ {
			if env[i] == '=' {
				envValues[env[:i]] = env[i+1:]
			}
		}
	}
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Ptr:
		elem := t.Elem()
		values := make(map[string]interface{})
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			argName, ok := field.Tag.Lookup("env")
			if !ok {
				continue
			}
			value, ok := envValues[argName]
			if !ok {
				value, ok = field.Tag.Lookup("default")
				if !ok {
					continue
				}
			}
			switch field.Type.Kind() {
			case reflect.String:
				values[field.Name] = value
			case reflect.Bool:
				values[field.Name], err = strconv.ParseBool(value)
				if err != nil {
					return nil, errors.Errorf("parse field %s as bool failed: %v", field.Name, err)
				}
			case reflect.Int:
				values[field.Name], err = strconv.Atoi(value)
				if err != nil {
					return nil, errors.Errorf("parse field %s as int failed: %v", field.Name, err)
				}
			default:
				return nil, errors.Errorf("field %s with type %s is unsupported", field.Name, field.Type.Kind())
			}
		}
		return values, nil
	default:
		return nil, errors.Errorf("need pointer")
	}
}
