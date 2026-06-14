// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"strconv"
)

// Business limits for the checkout service. These are centralized here so they
// can be located and changed in one place, and overridden per environment via
// the corresponding environment variables.
const (
	// defaultMaxOrderItems is the fallback maximum number of items allowed in a
	// single order when MAX_ORDER_ITEMS is not set.
	defaultMaxOrderItems = 50
)

// maxOrderItemsLimit returns the configured maximum number of items allowed in a
// single order. It reads the MAX_ORDER_ITEMS environment variable and falls back
// to defaultMaxOrderItems when unset or invalid.
func maxOrderItemsLimit() int {
	if v := os.Getenv("MAX_ORDER_ITEMS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	return defaultMaxOrderItems
}
