package main

import (
	"context"

	"github.com/jczz/shared/ci/dagger/internal/dagger"
)

type Shared struct{}

func (m *Shared) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"echo", stringArg})
}

func (m *Shared) DeployRemote(
	ctx context.Context,
	source *dagger.Directory,
	appName string,
	resourceGroup string,
	deploymentToken string,
) error {
	c := dag.Container().
		From("node:18-alpine").
		WithMountedDirectory("/app", source).
		WithWorkdir("/app").
		WithExec([]string{"npm", "install", "-g", "@azure/static-web-apps-cli"}).
		WithExec([]string{
			"swa", "deploy", ".",
			"--env", "production",
			"--app-name", appName,
			"--resource-group", resourceGroup,
			"--deployment-token", deploymentToken,
		})

	_, err := c.Sync(ctx)
	return err
}
