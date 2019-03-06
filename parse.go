// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"os"
)

// Parse env and flag to the given struct.
func Parse(data interface{}) error {
	argsValues, err := argsToValues(data, os.Args[1:])
	if err != nil {
		return err
	}
	envValues, err := envToValues(data, os.Environ())
	if err != nil {
		return err
	}
	return fill(data, mergeValues(argsValues, envValues))
}

func mergeValues(list ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, values := range list {
		for k, v := range values {
			result[k] = v
		}
	}
	return result
}
