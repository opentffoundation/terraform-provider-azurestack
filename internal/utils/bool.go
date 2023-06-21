// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package utils

// NormaliseNilableBool takes a pointer to a bool and returns a zero value or
// the real value if present
func NormaliseNilableBool(input *bool) bool {
	if input == nil {
		return false
	}

	return *input
}
