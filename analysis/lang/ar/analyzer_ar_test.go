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

package ar

import (
	"reflect"
	"testing"

	"github.com/appbaseio/bluge/analysis"
)

func TestArabicAnalyzer(t *testing.T) {
	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		{
			input: []byte("كبير"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("كبير"),
					PositionIncr: 1,
					Start:        0,
					End:          8,
				},
			},
		},
		// feminine marker
		{
			input: []byte("كبيرة"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("كبير"),
					PositionIncr: 1,
					Start:        0,
					End:          10,
				},
			},
		},
		{
			input: []byte("مشروب"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("مشروب"),
					PositionIncr: 1,
					Start:        0,
					End:          10,
				},
			},
		},
		// plural -at
		{
			input: []byte("مشروبات"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("مشروب"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		// plural -in
		{
			input: []byte("أمريكيين"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("امريك"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		// singular with bare alif
		{
			input: []byte("امريكي"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("امريك"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{
			input: []byte("كتاب"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("كتاب"),
					PositionIncr: 1,
					Start:        0,
					End:          8,
				},
			},
		},
		// definite article
		{
			input: []byte("الكتاب"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("كتاب"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{
			input: []byte("ما ملكت أيمانكم"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("ملكت"),
					PositionIncr: 2,
					Start:        5,
					End:          13,
				},
				&analysis.Token{
					Term:         []byte("ايمانكم"),
					PositionIncr: 1,
					Start:        14,
					End:          28,
				},
			},
		},
		// stopwords
		{
			input: []byte("الذين ملكت أيمانكم"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("ملكت"),
					PositionIncr: 2,
					Start:        11,
					End:          19,
				},
				&analysis.Token{
					Term:         []byte("ايمانكم"),
					PositionIncr: 1,
					Start:        20,
					End:          34,
				},
			},
		},
		// presentation form normalization
		{
			input: []byte("ﺍﻟﺴﻼﻢ"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("سلام"),
					PositionIncr: 1,
					Start:        0,
					End:          15,
				},
			},
		},
	}

	analyzer := Analyzer()
	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("expected %v, got %v", test.output, actual)
			t.Errorf("expected % x, got % x", test.output[0].Term, actual[0].Term)
		}
	}
}
