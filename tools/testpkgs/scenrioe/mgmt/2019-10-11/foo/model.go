// Copyright IBM Corp. 2014, 2021
// SPDX-License-Identifier: Apache-2.0

package foo

import "github.com/Azure/go-autorest/autorest"

// Gateway ...
type Gateway struct {
	autorest.Response `json:"-"`
	// Field1 ...
	Field1 *string
	// Field2 ...
	Field2 *int
}
