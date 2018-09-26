// Copyright 2018 The Prizem Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/prizem-io/api/v1/convert"
	"github.com/prizem-io/api/v1/proto"
)

func TestAttributes(t *testing.T) {
	var attributes proto.Attributes
	var actual map[string]interface{}
	attributes = convert.EncodeAttributes(nil)
	assert.Equal(t, 0, len(attributes.Words))
	actual = convert.DecodeAttributes(attributes)
	assert.NotNil(t, actual)
	assert.Equal(t, 0, len(actual))

	now := time.Now()

	m := map[string]interface{}{
		"01": "foo",
		"02": []string{"foo", "bar"},
		"03": 1,
		"04": int8(2),
		"05": int16(3),
		"06": int32(4),
		"07": int64(5),
		"08": int(6),
		"09": []int8{2, 3},
		"10": []int16{3, 4},
		"11": []int32{4, 5},
		"12": []int64{5, 6},
		"13": []int{6, 7},
		"14": float32(7),
		"15": float64(8),
		"16": []float32{6, 7},
		"17": []float64{7, 8},
		"18": true,
		"19": []bool{true, false},
		"20": now,
		"21": []time.Time{now.Add(time.Second), now.Add(time.Minute)},
		"22": time.Second,
		"23": []time.Duration{time.Second, time.Minute},
		"24": []byte("foo"),
		"25": map[string]interface{}{
			"foo": 1,
			"bar": 2,
		},
	}
	extected := map[string]interface{}{
		"01": "foo",
		"02": []string{"foo", "bar"},
		"03": int64(1),
		"04": int64(2),
		"05": int64(3),
		"06": int64(4),
		"07": int64(5),
		"08": int64(6),
		"09": []int64{2, 3},
		"10": []int64{3, 4},
		"11": []int64{4, 5},
		"12": []int64{5, 6},
		"13": []int64{6, 7},
		"14": float64(7),
		"15": float64(8),
		"16": []float64{6, 7},
		"17": []float64{7, 8},
		"18": true,
		"19": []bool{true, false},
		"20": now,
		"21": []time.Time{now.Add(time.Second), now.Add(time.Minute)},
		"22": time.Second,
		"23": []time.Duration{time.Second, time.Minute},
		"24": []byte("foo"),
		"25": map[string]interface{}{
			"foo": int64(1),
			"bar": int64(2),
		},
	}
	attributes = convert.EncodeAttributes(m)
	actual = convert.DecodeAttributes(attributes)
	assert.Equal(t, extected, actual)
	assert.NotNil(t, actual)
}
