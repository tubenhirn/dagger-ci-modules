package cimodules

import (
	"context"
	"strconv"
	"time"

	"dagger.io/dagger"
)

type RenovateOpts struct {
	Platform           string
	Autodiscover       bool
	AutodiscoverFilter string
	Repositories       string
	Env                map[string]string
	Secret             map[string]dagger.SecretID
	LogLevel           string
	Image              Image
}

var defaultRenovateImage = Image{
	Name: "renovate/renovate",
	//# renovate: datasource=docker depName=renovate/renovate versioning=docker
	Version: "38.135.1",
	Suffix:  "",
}

func renovate(ctx context.Context, client dagger.Client, opts RenovateOpts) error {
	image := client.Container().From(createImageString(defaultRenovateImage, opts.Image))

	// used to avoid dagger caching
	// we want this function to be executed every time we run it
	cacheHack := time.Now()

	renovate := image.WithEnvVariable("RENOVATE_PLATFORM", opts.Platform)
	renovate = renovate.WithEnvVariable("RENOVATE_EXTENDS", "github>whitesource/merge-confidence:beta")
	renovate = renovate.WithEnvVariable("RENOVATE_REQUIRE_CONFIG", "required")
	renovate = renovate.WithEnvVariable("RENOVATE_GIT_AUTHOR", "Renovate Bot <bot@renovateapp.com>")
	renovate = renovate.WithEnvVariable("RENOVATE_PIN_DIGEST", "true")
	renovate = renovate.WithEnvVariable("RENOVATE_DEPENDENCY_DASHBOARD", "false")
	renovate = renovate.WithEnvVariable("RENOVATE_LABELS", "renovate")
	renovate = renovate.WithEnvVariable("RENOVATE_AUTODISCOVER", strconv.FormatBool(opts.Autodiscover))
	renovate = renovate.WithEnvVariable("RENOVATE_AUTODISCOVER_FILTER", opts.AutodiscoverFilter)
	renovate = renovate.WithEnvVariable("RENOVATE_REPOSITORIES", opts.Repositories)
	renovate = renovate.WithEnvVariable("LOG_LEVEL", opts.LogLevel)
	// pass this value to avoid dagger caching
	// we want this container to be executed every time we run it
	renovate = renovate.WithEnvVariable("CACHE_HACK", cacheHack.String())

	// write env secrets - access-tokens etc.
	for key, val := range opts.Secret {
		renovate = renovate.WithSecretVariable(key, client.LoadSecretFromID(val))
	}

	// write dditional env variables
	for key, val := range opts.Env {
		renovate = renovate.WithEnvVariable(key, val)
	}

	_, err := renovate.WithExec([]string{}, dagger.ContainerWithExecOpts{}).Stdout(ctx)
	if err != nil {
		return err
	}

	return nil
}
