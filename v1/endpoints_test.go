// Copyright 2018 The Prizem Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api_test

import (
	"testing"

	"github.com/prizem-io/api/v1"
	"github.com/stretchr/testify/assert"
)

func TestEndpointsPrepare(t *testing.T) {
	endpoints := api.EndpointInfo{
		Nodes: []api.Node{
			{
				Services: []api.ServiceInstance{
					{
						Labels: []string{"foo", "bar"},
					},
				},
			},
		},
	}
	endpoints.Prepare()
	assert.Contains(t, endpoints.Nodes[0].Services[0].LabelSet, "foo")
	assert.Contains(t, endpoints.Nodes[0].Services[0].LabelSet, "bar")
}

func TestPortsFind(t *testing.T) {
	ports := api.Ports{
		{
			Port:     12345,
			Protocol: "HTTP1",
		},
		{
			Port:     12345,
			Protocol: "HTTP2",
			Secure:   true,
		},
	}

	port, ok := ports.Find("MQTT")
	assert.False(t, ok)

	port, ok = ports.Find("HTTP2")
	assert.True(t, ok)
	assert.Equal(t, int32(12345), port.Port)
	assert.Equal(t, "HTTP2", port.Protocol)
	assert.True(t, port.Secure)
}
