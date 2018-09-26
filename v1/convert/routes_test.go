// Copyright 2018 The Prizem Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/prizem-io/api/v1"
	"github.com/prizem-io/api/v1/convert"
)

func TestRoutes(t *testing.T) {
	assert.Nil(t, convert.EncodeServices(nil))
	assert.Nil(t, convert.EncodeVersion(nil))
	assert.Nil(t, convert.EncodeOperations(nil))
	assert.Nil(t, convert.EncodeRetry(nil))
	assert.Nil(t, convert.EncodeConfigurations(nil))
	assert.Nil(t, convert.EncodeHealthCheck(nil))

	assert.Nil(t, convert.DecodeServices(nil))
	assert.Nil(t, convert.DecodeVersion(nil))
	assert.Nil(t, convert.DecodeOperations(nil))
	assert.Nil(t, convert.DecodeRetry(nil))
	assert.Nil(t, convert.DecodeConfigurations(nil))
	assert.Nil(t, convert.DecodeHealthCheck(nil))

	services := []api.Service{
		{
			Name:        "service",
			Namespace:   "department.service",
			Description: "the service",
			Hostname:    "api.company.com",
			URIPrefix:   "/service",
			Version: &api.Version{
				VersionLocations: []string{"uri"},
				DefaultVersion:   "uri",
			},
			Authentication: "required",
			RoutingRules: api.RoutingRules{
				Selectors: []string{"v1.0.2"},
				RewriteRules: []api.Configuration{
					{
						Type: "test1",
						Config: map[string]interface{}{
							"foo": "bar",
						},
					},
				},
				Timeout: durationPtr(5 * time.Second),
				Retry: &api.Retry{
					Attempts:      3,
					PerTryTimeout: api.Duration(2 * time.Second),
				},
				Policies: []api.Configuration{
					{
						Type: "test2",
						Config: map[string]interface{}{
							"foo": "bar",
						},
					},
				},
			},
			Operations: []api.Operation{
				{
					Name:       "hello",
					Method:     "POST",
					URIPattern: "/hello",
					RoutingRules: api.RoutingRules{
						Selectors: []string{"v1.0.3"},
						RewriteRules: []api.Configuration{
							{
								Type: "test1",
								Config: map[string]interface{}{
									"foo": "bar",
								},
							},
						},
						Timeout: durationPtr(6 * time.Second),
						Retry: &api.Retry{
							Attempts:      3,
							PerTryTimeout: api.Duration(3 * time.Second),
						},
						Policies: []api.Configuration{
							{
								Type: "test2",
								Config: map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
			HealthCheck: &api.HealthCheck{
				Timeout:            api.Duration(5 * time.Second),
				Interval:           api.Duration(60 * time.Second),
				UnhealthyThreshold: 3,
				HealthyThreshold:   2,
				CheckType:          "http",
				CheckConfig: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
	}
	encoded := convert.EncodeServices(services)
	decoded := convert.DecodeServices(encoded)
	assert.Equal(t, services, decoded)
}

func durationPtr(d time.Duration) *api.Duration {
	ad := api.Duration(d)
	return &ad
}
