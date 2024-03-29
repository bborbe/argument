// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"context"
	"fmt"
	"log"
	"reflect"
)

// Print all configured arguments. Set display:"hidden" to hide or display:"length" to only print the arguments length.
func Print(ctx context.Context, data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		ef := e.Field(i)
		argName := t.Field(i).Tag.Get("display")
		if argName == "hidden" {
			continue
		}
		if argName == "length" {
			log.Printf("Argument: %s length %d", t.Field(i).Name, len(fmt.Sprintf("%v", ef.Interface())))
			continue
		}
		if ef.Kind() == reflect.Ptr || ef.Kind() == reflect.Interface {
			if ef.IsZero() {
				log.Printf("Argument: %s <nil>", t.Field(i).Name)
			} else {
				log.Printf("Argument: %s '%v'", t.Field(i).Name, ef.Elem())
			}
		} else {
			log.Printf("Argument: %s '%v'", t.Field(i).Name, ef.Interface())
		}
	}
	return nil
}
