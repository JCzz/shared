package main

import (
	"context"

	// IMPORTANT: this must match your module path in go.mod
	"github.com/jczz/shared/ci/dagger/internal/dagger"
)

type Shared struct{}

// Simple echo function for smoke testing
func (m *Shared) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"echo", stringArg})
}

// Deploy static site to Azure Static Web Apps using SWA CLI
func (m *Shared) DeploySwa(
	ctx context.Context,
	source *dagger.Directory, // usually ./public
	appName string,
	resourceGroup string,
	deploymentToken string,
) error {
	c := dag.Container().
		// you can swap this image later if you like
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
