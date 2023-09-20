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
	Version: "v4.0.11",
}

var defaultSemanticreleaseGitlabImage = Image{
	Name: "tubenhirn/semantic-release-gitlab",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-gitlab versioning=docker
	Version: "v4.0.9",
}

var defaultSemanticreleaseGitImage = Image{
	Name: "tubenhirn/semantic-release-git",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-git versioning=docker
	Version: "v4.0.9",
}

func semanticrelease(ctx context.Context, client dagger.Client, opts SemanticOpts) error {

	sourceDir := client.Host().Directory(opts.Source)

	var image string
	switch opts.Platform {
	case "github":
		image = createImageString(defaultSemanticreleaseGithubImage, opts.Image)
	case "gitlab":
		image = createImageString(defaultSemanticreleaseGitlabImage, opts.Image)
	case "git":
		image = createImageString(defaultSemanticreleaseGitImage, opts.Image)
	default:
		return errors.New("Platform net set.")
	}

	semanticreleaseImage := client.Container().From(image)

	semanticrelease := semanticreleaseImage.WithMountedDirectory("/src", sourceDir)
	semanticrelease = semanticrelease.WithWorkdir("/src")

	// write env secrets - access-tokens etc.
	for key, val := range opts.Secret {
		semanticrelease = semanticrelease.WithSecretVariable(key, client.Secret(val))
	}

	// write dditional env variables
	for key, val := range opts.Env {
		semanticrelease = semanticrelease.WithEnvVariable(key, val)
	}

	_, err := semanticrelease.WithExec(nil).Stdout(ctx)
	if err != nil {
		return err
	}

	return nil
}
