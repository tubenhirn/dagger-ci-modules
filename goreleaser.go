package cimodules

import (
	"context"

	"dagger.io/dagger"
)

type GoReleaserOpts struct {
	Source     string
	DryRun     bool `default:"false"`
	Snapshot   bool `default:"false"`
	RemoveDist bool `default:"false"`
	Env        map[string]string
	Secret     map[string]dagger.SecretID
	Image      Image
}

var defaultGoreleaserImage = Image{
	Name: "goreleaser/goreleaser",
	//# renovate: datasource=docker depName=goreleaser/goreleaser versioning=docker
	Version: "v1.24.0",
}

func release(ctx context.Context, client dagger.Client, opts GoReleaserOpts) error {

	commands := createFlags(opts)
	sourceDir := client.Host().Directory(opts.Source)

	goreleaser := client.Container().From(createImageString(defaultGoreleaserImage, opts.Image))
	goreleaser = goreleaser.WithMountedDirectory("/src", sourceDir)
	goreleaser = goreleaser.WithWorkdir("/src")

	// write env secrets - access-tokens etc.
	for key, val := range opts.Secret {
		goreleaser = goreleaser.WithSecretVariable(key, client.LoadSecretFromID(val))
	}

	// write env variables
	for key, val := range opts.Env {
		goreleaser = goreleaser.WithEnvVariable(key, val)
	}

	// set entrypoint
	goreleaser = goreleaser.WithEntrypoint([]string{"goreleaser"})

	// run the container
	goreleaser = goreleaser.WithExec(commands)

	// export the build artifacts to the host
	_, err := goreleaser.Directory("/src/dist").Export(ctx, opts.Source+"/dist")
	if err != nil {
		return err
	}

	return nil
}

func createFlags(opts GoReleaserOpts) []string {
	var flags []string
	if opts.DryRun {
		flags = append(flags, "--skip-publish")
		flags = append(flags, "--skip-announce")
	}
	if opts.Snapshot {
		flags = append(flags, "--snapshot")
	}
	if opts.RemoveDist {
		flags = append(flags, "--clean")
	}

	return flags
}
