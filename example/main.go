// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/bborbe/argument"
)

func main() {
	var data struct {
		Username string `arg:"username" default:"ben"`
		Password string `arg:"password"`
	}
	if err := argument.Parse(&data); err != nil {
		log.Fatalf("parse args failed: %v", err)
	}
	fmt.Printf("username %s, password %s\n", data.Username, data.Password)
}
