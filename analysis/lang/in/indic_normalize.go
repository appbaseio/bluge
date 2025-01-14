//  Copyright (c) 2020 Couchbase, Inc.
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

package in

import (
	"bytes"

	"github.com/appbaseio/bluge/analysis"
)

type IndicNormalizeFilter struct {
}

func NormalizeFilter() *IndicNormalizeFilter {
	return &IndicNormalizeFilter{}
}

func (s *IndicNormalizeFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	for _, token := range input {
		runes := bytes.Runes(token.Term)
		runes = normalize(runes)
		token.Term = analysis.BuildTermFromRunes(runes)
	}
	return input
}
