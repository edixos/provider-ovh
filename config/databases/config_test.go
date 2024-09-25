package databases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEndpointInformation(t *testing.T) {
	attrs := map[string]any{
		"endpoints": []map[string]any{
			{
				"uri":    "rediss://<username>:<password>@redis-526fe66a-od8a8be8e.database.cloud.ovh.net:20185",
				"domain": "redis-526fe66a-od8a8be8e.database.cloud.ovh.net",
				"port":   20185,
			},
		},
	}

	conn, err := addConnectionInfo(attrs)

	assert.Nil(t, err)
	assert.NotNil(t, conn)

	assert.Len(t, conn, 3)
	assert.Equal(t, []byte("rediss://<username>:<password>@redis-526fe66a-od8a8be8e.database.cloud.ovh.net:20185"), conn["endpoint_0_uri"])
	assert.Equal(t, []byte("redis-526fe66a-od8a8be8e.database.cloud.ovh.net"), conn["endpoint_0_domain"])
	assert.Equal(t, []byte("20185"), conn["endpoint_0_port"])
}

func TestAddEndpointInformationFailedDecode(t *testing.T) {
	attrs := map[string]any{
		"endpointss": []map[string]any{
			{
				"uri":    "rediss://<username>:<password>@redis-526fe66a-od8a8be8e.database.cloud.ovh.net:20185",
				"domain": "redis-526fe66a-od8a8be8e.database.cloud.ovh.net",
				"port":   20185,
			},
		},
	}

	conn, err := addConnectionInfo(attrs)

	assert.NotNil(t, err)
	assert.Nil(t, conn)
}
