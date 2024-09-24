package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func setupContainer(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "metrics-demo",       // Replace with your image name
		ExposedPorts: []string{"8080/tcp"}, // Replace with the port your app listens on
		WaitingFor:   wait.ForHTTP("/").WithStatusCodeMatcher(func(status int) bool { return status == 200 }),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := container.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatal(err)
	}

	return container, fmt.Sprintf("http://%s:%s", host, port.Port())
}

func TestRootRoute(t *testing.T) {
	container, url := setupContainer(t)
	defer container.Terminate(context.Background())

	resp, err := http.Get(url + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"message\":\"Hello, World!\"}", string(body)) // Replace with your expected response
}

func TestMetricsRoute(t *testing.T) {
	container, url := setupContainer(t)
	defer container.Terminate(context.Background())

	resp, err := http.Get(url + "/metrics")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotEmpty(t, string(body))
}
