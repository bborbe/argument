// Copyright (c) 2019 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bborbe/argument/v2"
)

type Username string

type Password string

type Active bool

func main() {
	ctx := context.Background()
	var data struct {
		Username Username `arg:"username" default:"ben"`
		Password Password `arg:"password" display:"length"`
		Active   *Active  `arg:"active"`
		Url      string   `arg:"url"`
	}
	if err := argument.Parse(ctx, &data); err != nil {
		log.Fatalf("parse args failed: %v", err)
	}
	fmt.Printf("%+v\n", data)
}
