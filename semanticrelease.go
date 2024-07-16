package cimodules

import (
	"context"
	"errors"

	"dagger.io/dagger"
)

type SemanticOpts struct {
	Source   string
	Platform string
	Env      map[string]string
	Secret   map[string]dagger.SecretID
	Image    Image
}

var defaultSemanticreleaseGithubImage = Image{
	Name: "tubenhirn/semantic-release-github",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-github versioning=docker
	Version: "v4.1.15",
}

var defaultSemanticreleaseGitlabImage = Image{
	Name: "tubenhirn/semantic-release-gitlab",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-gitlab versioning=docker
	Version: "v4.1.15",
}

var defaultSemanticreleaseGitImage = Image{
	Name: "tubenhirn/semantic-release-git",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-git versioning=docker
	Version: "v4.1.15",
}

var defaultSemanticreleaseAzueImage = Image{
	Name: "tubenhirn/semantic-release-azdo",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-azdo versioning=docker
	Version: "v4.1.15",
}

func semanticrelease(ctx context.Context, client dagger.Client, opts SemanticOpts) error {

	commands := []string{}
	sourceDir := client.Host().Directory(opts.Source)

	var image string
	switch opts.Platform {
	case "github":
		image = createImageString(defaultSemanticreleaseGithubImage, opts.Image)
	case "gitlab":
		image = createImageString(defaultSemanticreleaseGitlabImage, opts.Image)
	case "git":
		image = createImageString(defaultSemanticreleaseGitImage, opts.Image)
	case "azure":
		image = createImageString(defaultSemanticreleaseAzueImage, opts.Image)
	default:
		return errors.New("Platform net set.")
	}

	semanticrelease := client.Container().From(image)

	semanticrelease = semanticrelease.WithMountedDirectory("/src", sourceDir)
	semanticrelease = semanticrelease.WithWorkdir("/src")

	// write env secrets - access-tokens etc.
	for key, val := range opts.Secret {
		semanticrelease = semanticrelease.WithSecretVariable(key, client.LoadSecretFromID(val))
	}

	// write dditional env variables
	for key, val := range opts.Env {
		semanticrelease = semanticrelease.WithEnvVariable(key, val)
	}

	// set entrypoint
	semanticrelease = semanticrelease.WithEntrypoint([]string{"entrypoint.sh"})

	_, err := semanticrelease.WithExec(commands).Stdout(ctx)
	if err != nil {
		return err
	}

	return nil
}
