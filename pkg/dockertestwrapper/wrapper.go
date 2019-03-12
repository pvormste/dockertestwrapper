package dockertestwrapper

import (
	"github.com/ory/dockertest"
	"net/url"
	"strconv"
	"strings"
)

// DefaultContainerExpiresAfterSeconds tells docker the hard limit in seconds when the container should be purged
const DefaultContainerExpiresAfterSeconds uint = 1800

// AfterInitActionFunc is a function type which will be executed after container initialization
type AfterInitActionFunc func(dockerHost string, hostPort int) error

// WrapperParams contains all parameters needed to start a new custom container
type WrapperParams struct {
	ImageName           string
	ImageVersion        string
	EnvVariables        []string
	ContainerPort       string
	AfterInitActionFunc AfterInitActionFunc
}

// WrapperInstance holds all the information of the running container
type WrapperInstance struct {
	DockerHost string
	HostPort   int
	Pool       *dockertest.Pool
	Resource   *dockertest.Resource
}

// InitContainer starts a new container with the given parameters
func InitContainer(params WrapperParams) (instance *WrapperInstance, err error) {
	instance = &WrapperInstance{}
	instance.Pool, err = dockertest.NewPool("")
	if err != nil {
		return nil, err
	}

	instance.Resource, err = instance.Pool.Run(params.ImageName, params.ImageVersion, params.EnvVariables)
	if err != nil {
		return nil, err
	}

	if err := instance.Resource.Expire(DefaultContainerExpiresAfterSeconds); err != nil {
		return nil, err
	}

	if err := instance.determineDockerHost(); err != nil {
		return nil, err
	}

	if err := instance.determineHostPort(params.ContainerPort); err != nil {
		return nil, err
	}

	err = instance.Pool.Retry(func() error {
		return params.AfterInitActionFunc(instance.DockerHost, instance.HostPort)
	})
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// PurgeContainer purges the running container
func (w WrapperInstance) PurgeContainer() error {
	return w.Pool.Purge(w.Resource)
}

func (w *WrapperInstance) determineDockerHost() error {
	if strings.HasPrefix(w.Pool.Client.Endpoint(), "unix://") {
		w.DockerHost = "localhost"
		return nil
	}

	endpoint := w.Pool.Client.Endpoint()
	endpointUrl, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	w.DockerHost = endpointUrl.Hostname()
	return nil
}

func (w *WrapperInstance) determineHostPort(containerPort string) (err error) {
	stringPort := w.Resource.GetPort(containerPort)
	w.HostPort, err = strconv.Atoi(stringPort)
	if err != nil {
		return err
	}

	return nil
}
