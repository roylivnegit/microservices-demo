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
	"fmt"
	"math"

	pb "github.com/GoogleCloudPlatform/microservices-demo/src/shippingservice/genproto"
)

// nanosPerCent is the number of nanos (billionths of a currency unit) in one
// cent. Money.Nanos is expressed in 10^-9 units, so $0.01 == 10,000,000 nanos.
const nanosPerCent = 10_000_000

// Quote represents a currency value.
type Quote struct {
	Dollars uint32
	Cents   uint32
}

// String representation of the Quote.
func (q Quote) String() string {
	return fmt.Sprintf("$%d.%d", q.Dollars, q.Cents)
}

// ToMoney converts the Quote into a pb.Money value for the given currency code.
// Centralizing the cents-to-nanos conversion here avoids ad-hoc nanos scaling
// at call sites.
func (q Quote) ToMoney(currencyCode string) *pb.Money {
	return &pb.Money{
		CurrencyCode: currencyCode,
		Units:        int64(q.Dollars),
		Nanos:        int32(q.Cents * nanosPerCent),
	}
}

// CreateQuoteFromCount takes a number of items and returns a shipping quote.
func CreateQuoteFromCount(count int) Quote {
	if count == 0 {
		return CreateQuoteFromFloat(0)
	}
	return CreateQuoteFromFloat(8.99)
}

// CreateQuoteFromFloat takes a price represented as a float and creates a Price struct.
func CreateQuoteFromFloat(value float64) Quote {
	units, fraction := math.Modf(value)
	return Quote{
		uint32(units),
		uint32(math.Trunc(fraction * 100)),
	}
}