package dockertestwrapper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitPostgresContainer(t *testing.T) {
	t.Run("should start and purge a postgres 11 container successfully", func(t *testing.T) {
		wrapper, err := InitPostgresContainer(PostgresImageVersion11)
		assert.NoError(t, err)
		require.NotNil(t, wrapper)

		err = wrapper.PurgeContainer()
		assert.NoError(t, err)
	})

	t.Run("should start and purge a postgres 10 container successfully", func(t *testing.T) {
		wrapper, err := InitPostgresContainer(PostgresImageVersion10)
		assert.NoError(t, err)
		require.NotNil(t, wrapper)

		err = wrapper.PurgeContainer()
		assert.NoError(t, err)
	})
}
