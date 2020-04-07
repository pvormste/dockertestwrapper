package dockertestwrapper

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/ory/dockertest/v3"
)

// DefaultContainerExpiresAfterSeconds tells docker the hard limit in seconds when the container should be purged
const DefaultContainerExpiresAfterSeconds uint = 1800

// AfterInitActionFunc is a function type which will be executed after container initialization
type AfterInitActionFunc func(hostname string, port int) error

// WrapperParams contains all parameters needed to start a new custom container
type WrapperParams struct {
	ImageName           string
	ImageTag            string
	EnvVariables        []string
	ContainerPort       string
	AfterInitActionFunc AfterInitActionFunc
}

// WrapperInstance holds all the information of the running container
type WrapperInstance struct {
	Hostname   string
	Port       int
	DockerHost string // deprecated
	HostPort   int    // deprecated
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

	instance.Resource, err = instance.Pool.Run(params.ImageName, params.ImageTag, params.EnvVariables)
	if err != nil {
		return nil, err
	}

	if err := instance.Resource.Expire(DefaultContainerExpiresAfterSeconds); err != nil {
		return nil, err
	}

	if err := instance.determineHostname(); err != nil {
		return nil, err
	}

	if err := instance.determinePort(params.ContainerPort); err != nil {
		return nil, err
	}

	err = instance.Pool.Retry(func() error {
		return params.AfterInitActionFunc(instance.Hostname, instance.Port)
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

func (w *WrapperInstance) determineHostname() error {
	if strings.HasPrefix(w.Pool.Client.Endpoint(), "unix://") {
		w.Hostname = "localhost"
		w.DockerHost = w.Hostname // will be removed in a future update
		return nil
	}

	endpoint := w.Pool.Client.Endpoint()
	endpointUrl, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	w.Hostname = endpointUrl.Hostname()
	w.DockerHost = w.Hostname // will be removed in a future update
	return nil
}

func (w *WrapperInstance) determinePort(containerPort string) (err error) {
	stringPort := w.Resource.GetPort(containerPort)
	w.Port, err = strconv.Atoi(stringPort)
	if err != nil {
		return err
	}

	w.HostPort = w.Port // will be remove in a future update
	return nil
}
