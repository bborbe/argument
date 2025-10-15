// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"flag"
	"reflect"
	"strconv"
	"time"

	"github.com/bborbe/errors"
	libtime "github.com/bborbe/time"
)

// ParseArgs parses command-line arguments into the given struct using arg struct tags.
// See Parse() documentation for supported types and struct tag options.
//
// Parameters:
//   - ctx: Context for error handling
//   - data: Pointer to struct with arg tags
//   - args: Command-line arguments (typically os.Args[1:])
//
// Returns error if parsing fails or if default values are malformed.
func ParseArgs(ctx context.Context, data interface{}, args []string) error {
	values, err := argsToValues(ctx, data, args)
	if err != nil {
		return errors.Wrapf(ctx, err, "args to values failed")
	}
	if err := Fill(ctx, data, values); err != nil {
		return errors.Wrapf(ctx, err, "fill failed")
	}
	return nil
}

func handleCustomType(
	ctx context.Context,
	values map[string]interface{},
	tf reflect.StructField,
	ef reflect.Value,
	argName, defaultString string,
	found bool,
	usage string,
) (bool, error) {
	// Get the underlying type
	underlyingType := ef.Type()
	for underlyingType.Kind() == reflect.Ptr {
		underlyingType = underlyingType.Elem()
	}

	// Check if it's a named type (custom type) with an underlying primitive type
	if underlyingType.PkgPath() != "" && underlyingType.Kind() != reflect.Struct {
		switch underlyingType.Kind() {
		case reflect.String:
			values[tf.Name] = flag.CommandLine.String(argName, defaultString, usage)
			return true, nil
		case reflect.Bool:
			defaultValue, _ := strconv.ParseBool(defaultString)
			values[tf.Name] = flag.CommandLine.Bool(argName, defaultValue, usage)
			return true, nil
		case reflect.Int:
			defaultValue, _ := strconv.Atoi(defaultString)
			values[tf.Name] = flag.CommandLine.Int(argName, defaultValue, usage)
			return true, nil
		case reflect.Int64:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int64(argName, defaultValue, usage)
			return true, nil
		case reflect.Uint:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint(argName, uint(defaultValue), usage)
			return true, nil
		case reflect.Uint64:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint64(argName, defaultValue, usage)
			return true, nil
		case reflect.Int32:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int(argName, int(defaultValue), usage)
			return true, nil
		case reflect.Float64:
			defaultValue, _ := strconv.ParseFloat(defaultString, 64)
			values[tf.Name] = flag.CommandLine.Float64(argName, defaultValue, usage)
			return true, nil
		}
	}
	return false, nil
}

func argsToValues(
	ctx context.Context,
	data interface{},
	args []string,
) (map[string]interface{}, error) {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	values := make(map[string]interface{})
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("arg")
		if !ok {
			continue
		}
		defaultString, found := tf.Tag.Lookup("default")
		usage := tf.Tag.Get("usage")
		switch ef.Interface().(type) {
		case string:
			values[tf.Name] = flag.CommandLine.String(argName, defaultString, usage)
		case bool:
			defaultValue, _ := strconv.ParseBool(defaultString)
			values[tf.Name] = flag.CommandLine.Bool(argName, defaultValue, usage)
		case int:
			defaultValue, _ := strconv.Atoi(defaultString)
			values[tf.Name] = flag.CommandLine.Int(argName, defaultValue, usage)
		case int64:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int64(argName, defaultValue, usage)
		case uint:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint(argName, uint(defaultValue), usage)
		case uint64:
			defaultValue, _ := strconv.ParseUint(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Uint64(argName, defaultValue, usage)
		case int32:
			defaultValue, _ := strconv.ParseInt(defaultString, 10, 0)
			values[tf.Name] = flag.CommandLine.Int(argName, int(defaultValue), usage)
		case float64:
			defaultValue, _ := strconv.ParseFloat(defaultString, 64)
			values[tf.Name] = flag.CommandLine.Float64(argName, defaultValue, usage)
		case *float64:
			if found {
				defaultValue, _ := strconv.ParseFloat(defaultString, 64)
				values[tf.Name] = defaultValue
			}
			flag.CommandLine.Func(argName, usage, func(s string) error {
				if s == "" {
					return nil
				}
				v, err := strconv.ParseFloat(s, 64)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse float failed")
				}
				values[tf.Name] = v
				return nil
			})
		case time.Duration:
			if found {
				defaultValue, _ := libtime.ParseDuration(ctx, defaultString)
				if defaultValue != nil {
					values[tf.Name] = defaultValue.Duration()
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse duration failed")
				}
				values[tf.Name] = duration.Duration()
				return nil
			})
		case time.Time:
			if found {
				defaultValue, err := libtime.ParseTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				t, err := libtime.ParseTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse time failed")
				}
				values[tf.Name] = *t
				return nil
			})
		case *time.Time:
			if found {
				defaultValue, err := libtime.ParseTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				t, err := libtime.ParseTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse time failed")
				}
				values[tf.Name] = *t
				return nil
			})
		case *time.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = defaultValue.Duration()
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse duration failed")
				}
				values[tf.Name] = duration.Duration()
				return nil
			})
		case libtime.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse duration failed")
				}
				values[tf.Name] = *duration
				return nil
			})
		case *libtime.Duration:
			if found {
				defaultValue, err := libtime.ParseDuration(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				duration, err := libtime.ParseDuration(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse duration failed")
				}
				values[tf.Name] = *duration
				return nil
			})
		case libtime.DateTime:
			if found {
				defaultValue, err := libtime.ParseDateTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				dateTime, err := libtime.ParseDateTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse datetime failed")
				}
				values[tf.Name] = *dateTime
				return nil
			})
		case *libtime.DateTime:
			if found {
				defaultValue, err := libtime.ParseDateTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				dateTime, err := libtime.ParseDateTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse datetime failed")
				}
				values[tf.Name] = *dateTime
				return nil
			})
		case libtime.Date:
			if found {
				defaultValue, err := libtime.ParseDate(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				date, err := libtime.ParseDate(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse date failed")
				}
				values[tf.Name] = *date
				return nil
			})
		case *libtime.Date:
			if found {
				defaultValue, err := libtime.ParseDate(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				date, err := libtime.ParseDate(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse date failed")
				}
				values[tf.Name] = *date
				return nil
			})
		case libtime.UnixTime:
			if found {
				defaultValue, err := libtime.ParseUnixTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				unixTime, err := libtime.ParseUnixTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse unixtime failed")
				}
				values[tf.Name] = *unixTime
				return nil
			})
		case *libtime.UnixTime:
			if found {
				defaultValue, err := libtime.ParseUnixTime(ctx, defaultString)
				if err != nil {
					return nil, errors.Wrapf(ctx, err, "invalid default value %q for field %s", defaultString, tf.Name)
				}
				if defaultValue != nil {
					values[tf.Name] = *defaultValue
				}
			}
			flag.CommandLine.Func(argName, usage, func(value string) error {
				if value == "" {
					return nil
				}
				unixTime, err := libtime.ParseUnixTime(ctx, value)
				if err != nil {
					return errors.Wrapf(ctx, err, "parse unixtime failed")
				}
				values[tf.Name] = *unixTime
				return nil
			})
		default:
			// Check if it's a custom type with underlying primitive type
			if handled, err := handleCustomType(ctx, values, tf, ef, argName, defaultString, found, usage); handled {
				if err != nil {
					return nil, err
				}
			} else {
				return nil, errors.Errorf(ctx, "field %s with type %T is unsupported", tf.Name, ef.Interface())
			}
		}
	}
	if err := flag.CommandLine.Parse(args); err != nil {
		return nil, errors.Wrapf(ctx, err, "parse commandline failed")
	}
	return values, nil
}
