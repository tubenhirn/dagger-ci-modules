package goreleaser

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

type GoReleaserOpts struct {
	Source     string
	DryRun     bool `default:"false"`
	Snapshot   bool `default:"false"`
	RemoveDist bool `default:"false"`
	Env        map[string]string
	Secret     []string
}

type image struct {
	Name    string
	Version string
}

var goreleaserImage = image{
	Name: "goreleaser/goreleaser",
	//# renovate: datasource=docker depName=goreleaser/goreleaser versioning=docker
	Version: "v1.15.2",
}

func Release(ctx context.Context, opts GoReleaserOpts) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout), dagger.WithWorkdir("."))
	if err != nil {
		return err
	}
	defer client.Close()

	commands := createFlags(opts)
	sourceDir := client.Host().Directory(opts.Source)

	goreleaser := client.Container().From(createImageString(goreleaserImage))
	goreleaser = goreleaser.WithMountedDirectory("/src", sourceDir)
	goreleaser = goreleaser.WithWorkdir("/src")

	// write env secrets
	for _, ele := range opts.Secret {
		secret, err := client.Host().EnvVariable(ele).Secret().ID(ctx)
		if err != nil {
			return err
		}
		goreleaser = goreleaser.WithSecretVariable(ele, client.Secret(secret))
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
	_, err =  goreleaser.Directory("/src/dist").Export(ctx, opts.Source + "/dist")
	if err != nil{
		return err
	}

	return nil
}

func createFlags(opts GoReleaserOpts) []string {
	var flags []string
	// flags = append(flags, "release")
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

func createImageString(img image) string {
	return fmt.Sprintf("%s:%s", img.Name, img.Version)
}
