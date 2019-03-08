package argument

import (
	"fmt"
	"reflect"

	"github.com/golang/glog"
)

func print(data interface{}) error {
	e := reflect.ValueOf(data).Elem()
	t := e.Type()
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		argName := t.Field(i).Tag.Get("display")
		if argName == "hidden" {
			continue
		}
		if argName == "length" {
			glog.V(0).Infof("Argument: %s length %d", t.Field(i).Name, len(fmt.Sprintf("%v", f.Interface())))
			continue
		}
		glog.V(0).Infof("Argument: %s '%v'", t.Field(i).Name, f.Interface())
	}
	return nil
}
