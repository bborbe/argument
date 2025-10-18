// Copyright (c) 2025 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package argument

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/bborbe/errors"
)

// Fill populates the given struct with values from the provided map using JSON marshaling.
// It encodes the map to JSON and then decodes it back into the struct, allowing for
// flexible type conversion and nested structure population.
//
// Parameters:
//   - ctx: Context for error handling
//   - data: Pointer to struct to populate
//   - values: Map of field names to values
//
// Returns error if JSON encoding or decoding fails.
func Fill(ctx context.Context, data interface{}, values map[string]interface{}) error {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(values); err != nil {
		return errors.Wrap(ctx, err, "encode json failed")
	}
	if err := json.NewDecoder(buf).Decode(data); err != nil {
		return errors.Wrap(ctx, err, "decode json failed")
	}
	return nil
}
