//  Copyright (c) 2020 The Bluge Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggregations

import "github.com/appbaseio/bluge/search"

var staticCount = []float64{1}

type countingSource struct{}

func (*countingSource) Fields() []string {
	return nil
}

func (*countingSource) Numbers(_ *search.DocumentMatch) []float64 {
	return staticCount
}

var countSource = &countingSource{}

func CountMatches() *SingleValueMetric {
	return Sum(countSource)
}
