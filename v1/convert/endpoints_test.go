// Copyright 2018 The Prizem Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package convert_test

import (
	"net"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	"github.com/prizem-io/api/v1"
	"github.com/prizem-io/api/v1/convert"
)

func TestEndpoints(t *testing.T) {
	assert.Nil(t, convert.EncodeNodes(nil))
	assert.Nil(t, convert.EncodeServiceInstances(nil))
	assert.Nil(t, convert.EncodeContainer(nil))
	assert.Nil(t, convert.EncodePorts(nil))

	assert.Nil(t, convert.DecodeNodes(nil))
	assert.Nil(t, convert.DecodeServiceInstances(nil))
	assert.Nil(t, convert.DecodeContainer(nil))
	assert.Nil(t, convert.DecodePorts(nil))

	id := uuid.NewV4()
	nodes := []api.Node{
		{
			ID:         id,
			Address:    net.ParseIP("127.0.0.1"),
			Datacenter: "us-east",
			Geography:  "east",
			Metadata: api.Metadata{
				"foo": "bar",
			},
			Services: []api.ServiceInstance{
				{
					ID:        id,
					Service:   "service",
					Name:      "service01",
					Namespace: "department.service",
					Principal: "service-account",
					Owner:     "department",
					Container: &api.Container{
						Name:  "service",
						Image: "repo.org/service:1.0.1",
					},
					Ports: api.Ports{
						{
							Port:     443,
							Protocol: "HTTP2",
							Secure:   true,
						},
					},
					Metadata: api.Metadata{
						"version": "1.0.1",
					},
					Labels: []string{"foo", "bar"},
				},
			},
		},
	}
	encoded := convert.EncodeNodes(nodes)
	decoded := convert.DecodeNodes(encoded)
	assert.Equal(t, nodes, decoded)
}
