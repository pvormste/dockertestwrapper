package dockertestwrapper

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ===========================
// wrapper.determineHostname()
// ===========================

var wrapperInstanceDetermineHostnameTests = []struct {
	it                 string
	inputDockerAddress string
	doAssertions       func(t *testing.T, actualInstance WrapperInstance, actualErr error)
}{
	{
		it:                 "should return localhost when starting with unix://",
		inputDockerAddress: "unix:///var/run/docker.sock",
		doAssertions: func(t *testing.T, actualInstance WrapperInstance, actualErr error) {
			assert.NoError(t, actualErr)
			assert.Equal(t, "localhost", actualInstance.Hostname)
		},
	},
	{
		it:                 "should return the host from pool when not starting with unix://",
		inputDockerAddress: "http://docker:2375",
		doAssertions: func(t *testing.T, actualInstance WrapperInstance, actualErr error) {
			assert.NoError(t, actualErr)
			assert.Equal(t, "docker", actualInstance.Hostname)
		},
	},
}

func TestWrapperInstance_DetermineHostname(t *testing.T) {
	for _, test := range wrapperInstanceDetermineHostnameTests {
		test := test
		t.Run(test.it, func(t *testing.T) {
			client, err := docker.NewClient(test.inputDockerAddress)
			require.NoError(t, err)

			actualWrapper := WrapperInstance{
				Pool: &dockertest.Pool{
					Client: client,
				},
			}
			actualErr := actualWrapper.determineHostname()
			test.doAssertions(t, actualWrapper, actualErr)
		})
	}
}

// ===========================
// wrapper.determinePort()
// ===========================

var wrapperInstanceDeterminePortTests = []struct {
	it                 string
	inputContainerPort string
	internalHostPort   string
	doAssertions       func(t *testing.T, instance WrapperInstance, actualErr error)
}{
	{
		it:                 "should return UnassignedPort when no container port is provided",
		inputContainerPort: "",
		internalHostPort:   "5432/tcp",
		doAssertions: func(t *testing.T, instance WrapperInstance, actualErr error) {
			assert.NoError(t, actualErr)
			assert.Equal(t, UnassignedPort, instance.Port)
		},
	},
	{
		it:                 "should return error when container port is an invalid port",
		inputContainerPort: "5432/tcp",
		internalHostPort:   "5432/tcp",
		doAssertions: func(t *testing.T, instance WrapperInstance, actualErr error) {
			assert.Error(t, actualErr)
			assert.Equal(t, 0, instance.Port)
		},
	},
	{
		it:                 "should determine successfully the port on host",
		inputContainerPort: "5432/tcp",
		internalHostPort:   "5432",
		doAssertions: func(t *testing.T, instance WrapperInstance, actualErr error) {
			assert.NoError(t, actualErr)
			assert.Equal(t, 5432, instance.Port)
		},
	},
}

func TestWrapperInstance_DeterminePort(t *testing.T) {
	for _, test := range wrapperInstanceDeterminePortTests {
		test := test
		t.Run(test.it, func(t *testing.T) {
			mockedPortBinding := map[docker.Port][]docker.PortBinding{}
			mockedPortBinding[docker.Port(test.inputContainerPort)] = []docker.PortBinding{
				{
					HostPort: test.internalHostPort,
				},
			}

			mockedResource := dockertest.Resource{
				Container: &docker.Container{
					NetworkSettings: &docker.NetworkSettings{
						Ports: mockedPortBinding,
					},
				},
			}

			actualWrapper := WrapperInstance{
				Resource: &mockedResource,
			}

			actualErr := actualWrapper.determinePort(test.inputContainerPort)
			test.doAssertions(t, actualWrapper, actualErr)
		})
	}
}
