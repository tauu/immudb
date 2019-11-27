/*
Copyright 2019 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"fmt"
	"os"
	"strconv"
)

type Options struct {
	Address            string
	Port               int
	DialRetries        int
	HealthCheckRetries int
}

func DefaultOptions() Options {
	return Options{
		Address:            "127.0.0.1",
		Port:               8080,
		DialRetries:        5,
		HealthCheckRetries: 5,
	}
}

func (o Options) WithAddress(address string) Options {
	o.Address = address
	return o
}

func (o Options) WithPort(port int) Options {
	o.Port = port
	return o
}

func (o Options) Bind() string {
	return o.Address + ":" + strconv.Itoa(o.Port)
}

func (o Options) String() string {
	return fmt.Sprintf("{address:%v port:%d}", o.Address, o.Port)
}

func (o Options) FromEnvironment() Options {
	address := os.Getenv("IMMUDB_ADDRESS")
	if address != "" {
		o.Address = address
	}
	port := os.Getenv("IMMUDB_PORT")
	if parsedPort, err := strconv.Atoi(port); err == nil {
		o.Port = parsedPort
	}
	dialRetries := os.Getenv("IMMUDB_DIAL_RETRIES")
	if parsedDialRetries, err := strconv.Atoi(dialRetries); err == nil {
		o.DialRetries = parsedDialRetries
	}
	healthCheckRetries := os.Getenv("IMMUDB_HEALTH_CHECK_RETRIES")
	if parsedHealthCheckRetries, err := strconv.Atoi(healthCheckRetries); err == nil {
		o.HealthCheckRetries = parsedHealthCheckRetries
	}
	return o
}
