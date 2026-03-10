// Copyright IBM Corp. 2014, 2021
// SPDX-License-Identifier: Apache-2.0

package utils

import "strings"

// NormalizePath normalizes the path by replacing \ with /
func NormalizePath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}
