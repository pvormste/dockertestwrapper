package dockertestwrapper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitPostgres11Container(t *testing.T) {
	t.Run("should start and purge successfully", func(t *testing.T) {
		wrapper, err := InitPostgres11Container()
		assert.NoError(t, err)
		require.NotNil(t, wrapper)

		err = wrapper.PurgeContainer()
		assert.NoError(t, err)
	})
}

func TestInitPostgres10Container(t *testing.T) {
	t.Run("should start and purge successfully", func(t *testing.T) {
		wrapper, err := InitPostgres10Container()
		assert.NoError(t, err)
		require.NotNil(t, wrapper)

		err = wrapper.PurgeContainer()
		assert.NoError(t, err)
	})
}
