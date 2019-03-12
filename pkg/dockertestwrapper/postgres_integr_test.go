package dockertestwrapper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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
