# dockertestwrapper [![Build Status](https://travis-ci.org/pvormste/dockertestwrapper.svg?branch=master)](https://travis-ci.org/pvormste/dockertestwrapper)

As the name suggests dockertestwrapper is a wrapper for orys awesome [dockertest library](https://github.com/ory/dockertest).  
It provides an easy to use api to be used in your integration tests.

## Usage

### Postgres

#### Postgres 11

```go
dockertesthelper, err := dockertestwrapper.InitPostgres11Container()
if err != nil {
	// ...
}


if err := dockertesthelper.PurgeContainer(); err != nil {
	// ...
}
```

#### Postgres 10

```go
dockertesthelper, err := dockertestwrapper.InitPostgres10Container()
if err != nil {
	// ...
}


if err := dockertesthelper.PurgeContainer(); err != nil {
	// ...
}
```

#### Postgres Custom Version

```go
dockertesthelper, err := dockertestwrapper.InitPostgresContainer("9.6")
if err != nil {
	// ...
}


if err := dockertesthelper.PurgeContainer(); err != nil {
	// ...
}
```

### Custom Container 

```go
params := dockertestwrapper.WrapperParams{
	ImageName: "golang",
	ImageVersion: "1.12",
	EnvVariables: []string{},
	ContainerPort: 80,
	AfterInitActionFunc: func(dockerHost string, hostPort int) error {
		// Start a webserver or something
		return nil
	},
}

dockertesthelper, err := dockertestwrapper.InitContainer(params)
if err != nil {
	// ...
}


if err := dockertesthelper.PurgeContainer(); err != nil {
	// ...
}
```

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