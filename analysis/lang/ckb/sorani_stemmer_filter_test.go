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

package ckb

import (
	"reflect"
	"testing"

	"github.com/appbaseio/bluge/analysis/tokenizer"

	"github.com/appbaseio/bluge/analysis"
)

func TestSoraniStemmerFilter(t *testing.T) {
	// in order to match the lucene tests
	// we will test with an analyzer, not just the stemmer
	analyzer := analysis.Analyzer{
		Tokenizer: tokenizer.NewSingleTokenTokenizer(),
		TokenFilters: []analysis.TokenFilter{
			NormalizeFilter(),
			StemmerFilter(),
		},
	}

	tests := []struct {
		input  []byte
		output analysis.TokenStream
	}{
		{ // -ek
			input: []byte("پیاوێک"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{ // -yek
			input: []byte("دەرگایەک"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // -aka
			input: []byte("پیاوەكە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -ka
			input: []byte("دەرگاكە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -a
			input: []byte("کتاویە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("کتاوی"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{ // -ya
			input: []byte("دەرگایە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -An
			input: []byte("پیاوان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{ // -yAn
			input: []byte("دەرگایان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // -akAn
			input: []byte("پیاوەکان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // -kAn
			input: []byte("دەرگاکان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // -Ana
			input: []byte("پیاوانە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پیاو"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -yAna
			input: []byte("دەرگایانە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دەرگا"),
					PositionIncr: 1,
					Start:        0,
					End:          18,
				},
			},
		},
		{ // Ezafe singular
			input: []byte("هۆتیلی"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("هۆتیل"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{ // Ezafe indefinite
			input: []byte("هۆتیلێکی"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("هۆتیل"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // Ezafe plural
			input: []byte("هۆتیلانی"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("هۆتیل"),
					PositionIncr: 1,
					Start:        0,
					End:          16,
				},
			},
		},
		{ // -awa
			input: []byte("دوورەوە"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("دوور"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -dA
			input: []byte("نیوەشەودا"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("نیوەشەو"),
					PositionIncr: 1,
					Start:        0,
					End:          18,
				},
			},
		},
		{ // -A
			input: []byte("سۆرانا"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("سۆران"),
					PositionIncr: 1,
					Start:        0,
					End:          12,
				},
			},
		},
		{ // -mAn
			input: []byte("پارەمان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پارە"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -tAn
			input: []byte("پارەتان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پارە"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // -yAn
			input: []byte("پارەیان"),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte("پارە"),
					PositionIncr: 1,
					Start:        0,
					End:          14,
				},
			},
		},
		{ // empty
			input: []byte(""),
			output: analysis.TokenStream{
				&analysis.Token{
					Term:         []byte(""),
					PositionIncr: 1,
					Start:        0,
					End:          0,
				},
			},
		},
	}

	for _, test := range tests {
		actual := analyzer.Analyze(test.input)
		if !reflect.DeepEqual(actual, test.output) {
			t.Errorf("for input %s(% x)", test.input, test.input)
			t.Errorf("\texpected:")
			for _, token := range test.output {
				t.Errorf("\t\t%v %s(% x)", token, token.Term, token.Term)
			}
			t.Errorf("\tactual:")
			for _, token := range actual {
				t.Errorf("\t\t%v %s(% x)", token, token.Term, token.Term)
			}
		}
	}
}
