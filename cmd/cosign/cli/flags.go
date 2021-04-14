//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import "reflect"

// oneOf ensures that only one of the supplied interfaces is set to a non-zero value.
func oneOf(args ...interface{}) bool {
	foundOne := false
	for _, arg := range args {
		if !reflect.ValueOf(arg).IsZero() {
			if foundOne {
				return false
			}
			foundOne = true
		}
	}
	return foundOne
}

// allOf ensures that all of the supplied interfaces are set to a non-zero value.
func allOf(args ...interface{}) bool {
	foundAll := false
	for _, arg := range args {
		if reflect.ValueOf(arg).IsZero() {
			return false
		}
		foundAll = true
	}
	return foundAll
}
