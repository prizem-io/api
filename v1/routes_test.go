// Copyright 2018 The Prizem Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/prizem-io/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestDurationEncode(t *testing.T) {
	var d api.Duration
	var err error

	err = json.Unmarshal([]byte(`"asdf"`), &d)
	assert.NotNil(t, err)

	err = json.Unmarshal([]byte(`null`), &d)
	assert.Nil(t, err)
	assert.Equal(t, api.Duration(0), d)

	err = json.Unmarshal([]byte(`"3s"`), &d)
	assert.Nil(t, err)
	assert.Equal(t, api.Duration(3*time.Second), d)
}

func TestDurationDecode(t *testing.T) {
	d := api.Duration(3 * time.Second)
	data, err := json.Marshal(&d)
	assert.Nil(t, err)
	assert.Equal(t, []byte(`"3s"`), data)
}
