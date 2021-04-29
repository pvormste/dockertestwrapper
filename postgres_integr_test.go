package dockertestwrapper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var postgresImageTags = map[string]string{
	"13":  "13-alpine",
	"12":  "12-alpine",
	"11":  "11-alpine",
	"10":  "10-alpine",
	"9.6": "9.6-alpine",
}

func TestInitPostgresContainer(t *testing.T) {
	for postgresVersion, postgresTag := range postgresImageTags {
		it := fmt.Sprintf("should start and purge a postgres %s container successfully", postgresVersion)
		t.Run(it, func(t *testing.T) {
			wrapper, err := InitPostgresContainer(postgresTag)
			assert.NoError(t, err)
			require.NotNil(t, wrapper)

			err = wrapper.PurgeContainer()
			assert.NoError(t, err)
		})
	}
}
