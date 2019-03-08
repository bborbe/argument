package argument

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

func validateRequired(data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		tf := t.Field(i)
		ef := e.Field(i)
		argName, ok := tf.Tag.Lookup("required")
		if !ok || argName != "true" {
			continue
		}
		switch ef.Kind() {
		case reflect.String:
			if ef.Interface() == "" {
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
				return errors.New(buf.String())
			}
		}
	}
	return nil
}
