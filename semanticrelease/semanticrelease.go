package semanticrelease

import (
	"context"
	"errors"
	"fmt"

	"dagger.io/dagger"
)

type SemanticOpts struct {
	Source   string
	Platform string
	Env      map[string]string
	Secret   map[string]dagger.SecretID
}

type image struct {
	Name    string
	Version string
}

var semanticreleaseGithubImage = image{
	Name: "tubenhirn/semantic-release-github",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-github versioning=docker
	Version: "v3.0.0",
}

var semanticreleaseGitlabImage = image{
	Name: "tubenhirn/semantic-release-gitlab",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-gitlab versioning=docker
	Version: "v3.0.1",
}

var semanticreleaseGitImage = image{
	Name: "tubenhirn/semantic-release-git",
	//# renovate: datasource=docker depName=tubenhirn/semantic-release-git versioning=docker
	Version: "v4.0.0",
}

func Semanticrelease(ctx context.Context, client dagger.Client, opts SemanticOpts) error {

	sourceDir := client.Host().Directory(opts.Source)

	var image string
	switch opts.Platform {
	case "github":
		image = createImageString(semanticreleaseGithubImage)
	case "gitlab":
		image = createImageString(semanticreleaseGitlabImage)
	case "git":
		image = createImageString(semanticreleaseGitImage)
	default:
		return errors.New("Platform net set.")
	}

	semanticrelease := client.Container().From(image)

	// write env secrets - access-tokens etc.
	for key, val := range opts.Secret {
		semanticrelease = semanticrelease.WithSecretVariable(key, client.Secret(val))
	}

	// write dditional env variables
	for key, val := range opts.Env {
		semanticrelease = semanticrelease.WithEnvVariable(key, val)
	}

	semanticrelease = semanticrelease.WithMountedDirectory("/src", sourceDir)
	semanticrelease = semanticrelease.WithWorkdir("/src")

	_, err := semanticrelease.Exec().Stdout(ctx)
	if err != nil {
		return err
	}

	return nil
}

func createImageString(img image) string {
	return fmt.Sprintf("%s:%s", img.Name, img.Version)
}