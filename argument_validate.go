// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/bborbe/errors"
)

func handleCustomTypeValidation(
	ctx context.Context,
	tf reflect.StructField,
	ef reflect.Value,
	createError func() error,
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
			// For custom string types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Bool:
			// Bool types are never considered "empty" for required validation
			return true, nil
		case reflect.Int, reflect.Int32, reflect.Int64:
			// For custom int types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Uint, reflect.Uint64:
			// For custom uint types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		case reflect.Float64:
			// For custom float64 types, check if value equals zero value of underlying type
			zeroValue := reflect.Zero(underlyingType).Interface()
			if ef.Interface() == zeroValue {
				return true, createError()
			}
			return true, nil
		}
	}
	return false, nil
}

// ValidateRequired fields are set and returns an error if not.
func ValidateRequired(ctx context.Context, data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("required")
		if !ok || argName != "true" {
			continue
		}
		createError := func() error {
			buf := bytes.NewBufferString("Required field empty, ")
			argName, argOk := tf.Tag.Lookup("arg")
			if argOk {
				fmt.Fprintf(buf, "define parameter %s", argName)
			}
			envName, envOk := tf.Tag.Lookup("env")
			if envOk {
				if argOk {
					fmt.Fprintf(buf, " or ")
				}
				fmt.Fprintf(buf, "define env %s", envName)
			}
			return errors.New(ctx, buf.String())
		}
		switch ef.Interface().(type) {
		case string:
			var empty string
			if empty == ef.Interface() {
				return createError()
			}
		case bool:
		case int:
			var empty int
			if empty == ef.Interface() {
				return createError()
			}
		case int64:
			var empty int64
			if empty == ef.Interface() {
				return createError()
			}
		case uint:
			var empty uint
			if empty == ef.Interface() {
				return createError()
			}
		case uint64:
			var empty uint64
			if empty == ef.Interface() {
				return createError()
			}
		case int32:
			var empty int32
			if empty == ef.Interface() {
				return createError()
			}
		case float64:
			var empty float64
			if empty == ef.Interface() {
				return createError()
			}
		case *float64:
			var empty *float64
			if empty == ef.Interface() {
				return createError()
			}
		case time.Duration:
			var empty time.Duration
			if empty == ef.Interface() {
				return createError()
			}
		default:
			// Handle slices
			if ef.Kind() == reflect.Slice {
				if ef.Len() == 0 {
					return createError()
				}
			} else {
				// Check if it's a custom type with underlying primitive type
				if handled, err := handleCustomTypeValidation(ctx, tf, ef, createError); handled {
					if err != nil {
						return err
					}
				} else {
					return errors.Errorf(ctx, "field %s with type %T is unsupported", tf.Name, ef.Interface())
				}
			}
		}
	}
	return nil
}
