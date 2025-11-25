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

func (m *Shared) Deploy(
	ctx context.Context,
	source *dagger.Directory,
	appName string,
	resourceGroup string,
	deploymentToken string,
) error {
	c := dag.Container().
		// Use Debian-based image so StaticSitesClient can run
		From("node:20-bullseye").
		// Optional but nice: ensure basic libs are there
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{
			"apt-get", "install", "-y",
			"libicu-dev", "ca-certificates", "curl",
		}).
		// Mount the source (./public) as /app
		WithMountedDirectory("/app", source).
		WithWorkdir("/app").
		// Let SWA CLI also see the token via env
		WithEnvVariable("SWA_CLI_DEPLOYMENT_TOKEN", deploymentToken).
		// Install SWA CLI
		WithExec([]string{"npm", "install", "-g", "@azure/static-web-apps-cli"}).
		// (Optional) tiny debug to see what we're deploying
		// .WithExec([]string{"sh", "-c", "echo '=== ls -R . ==='; ls -R ."}).
		// Run the actual deploy
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
