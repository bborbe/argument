// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"flag"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// ParseArgs parse os.Args[1:] into data.
func ParseArgs(data interface{}, args []string) error {
	values, err := argsToValues(data, args)
	if err != nil {
		return err
	}
	return fill(data, values)
}

func argsToValues(data interface{}, args []string) (map[string]interface{}, error) {
	t := reflect.TypeOf(data)
	switch t.Kind() {
	case reflect.Ptr:
		elem := t.Elem()
		values := make(map[string]interface{})
		flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			argName, ok := field.Tag.Lookup("arg")
			if !ok {
				continue
			}
			defaultString := field.Tag.Get("default")
			usage := field.Tag.Get("usage")
			switch field.Type.Kind() {
			case reflect.String:
				values[field.Name] = flagSet.String(argName, defaultString, usage)
			case reflect.Bool:
				defaultValue, _ := strconv.ParseBool(defaultString)
				values[field.Name] = flagSet.Bool(argName, defaultValue, usage)
			case reflect.Int:
				defaultValue, _ := strconv.Atoi(defaultString)
				values[field.Name] = flagSet.Int(argName, defaultValue, usage)
			default:
				return nil, errors.Errorf("field %s with type %s is unsupported", field.Name, field.Type.Kind())
			}
		}
		if err := flagSet.Parse(args); err != nil {
			return nil, err
		}
		return values, nil
	default:
		return nil, errors.Errorf("need pointer")
	}
}
