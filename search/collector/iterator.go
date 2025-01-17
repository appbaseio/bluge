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

package collector

import (
	"github.com/appbaseio/bluge/search"
)

type TopNIterator struct {
	results search.DocumentMatchCollection
	bucket  *search.Bucket
	index   int
	err     error
}

func (i *TopNIterator) Next() (*search.DocumentMatch, error) {
	if i.err != nil {
		return nil, i.err
	}
	if i.index < len(i.results) {
		rv := i.results[i.index]
		i.index++
		return rv, nil
	}
	return nil, nil
}

func (i *TopNIterator) Aggregations() *search.Bucket {
	return i.bucket
}
