# dockertestwrapper [![Build Status](https://travis-ci.org/pvormste/dockertestwrapper.svg?branch=master)](https://travis-ci.org/pvormste/dockertestwrapper)

As the name suggests dockertestwrapper is a wrapper for [orys awesome dockertest library](https://github.com/ory/dockertest).  
It provides an easy to use api to be used in your integration tests.

## Install

```go
go get -u github.com/pvormste/dockertestwrapper/v2
```

## Usage

### Postgres

The postgres helper function can be used to start a postgres container with default credentials/connection details.
Timezone defaults to `TZ=UTC`.

Although it should be possible to start any postgres version with the `InitPostgresContainer` function, following
images are covered by tests in this repository:
  - postgres:13-alpine
  - postgres:12-alpine
  - postgres:11-alpine
  - postgres:10-alpine
  - postgres:9.6-alpine

#### Connection details

| name | value |
| -----| ----- |
| host | `postgresContainer.Hostname` |
| port | `postgresContainer.Port` |
| user | postgres |
| password | postgres |
| database | postgres |

#### Start and purge a postgres container

```go
postgresContainer, err := dockertestwrapper.InitPostgresContainer("9.6")
if err != nil {
	// ...
}


if err := postgresContainer.PurgeContainer(); err != nil {
	// ...
}
```

### Custom Container 

You can start any container with this library, just populate the `WrapperParams` struct and pass it to the `InitContainer` function.  
The Hostname and Port will then be available via the fields `container.Hostname` and `container.Port`.

#### WrapperParams

| field | descritpion | example |
| ----- | ----------- | ------- |
| ImageName | name of the image | `"mysql`" |
| ImageTag | tag of the image | `"5.7"` |
| EnvVariables | env variables to be passed to container | `[]string{"MYSQL_ROOT_PASSWORD=mysql"}` |
| ContainerPort | exported port on container | `"3306/tcp"` |
| AfterInitActionFunc | function which will be executed after container initialization | see postgres.go for an example |

#### Example

```go
params := dockertestwrapper.WrapperParams{
	ImageName: "golang",
	ImageTag: "1.12",
	EnvVariables: []string{},
	ContainerPort: "80/tcp",
	AfterInitActionFunc: func(dockerHost string, hostPort int) error {
		// Start a webserver or something
		return nil
	},
}

customContainer, err := dockertestwrapper.InitContainer(params)
if err != nil {
	// ...
}


hostname := customContainer.Hostname
port := customContainer.Port


if err := customContainer.PurgeContainer(); err != nil {
	// ...
}
```

## Using with Gitlab CI

If you want to use this library with Gitlab CI you may be interested in my fork of the
docker image: [pvormste/docker-go](https://github.com/pvormste/docker)  
You can also find it on the DockerHub: [pvormste/docker-go](https://hub.docker.com/r/pvormste/docker-go)


## License

MIT License

Copyright (c) 2019 Patric Vormstein

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.