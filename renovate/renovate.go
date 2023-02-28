package renovate

import (
	"context"
	"fmt"
	"os"
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
	Secret             []string
	LogLevel           string
}

type image struct {
	Name    string
	Version string
}

var renovateImage = image{
	Name: "renovate/renovate",
	//# renovate: datasource=docker depName=renovate/renovate versioning=docker
	Version: "34.154.0",
}

func Renovate(ctx context.Context, opts RenovateOpts) error {
	// used to avoid dagger caching
	// we want this function to be executed every time we run it
	cacheHack := time.Now()
	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}

	defer client.Close()

	renovate := client.Container().From(createImageString(renovateImage))

	// write env secrets - access-tokens etc.
	for _, ele := range opts.Secret {
		secret, err := client.Host().EnvVariable(ele).Secret().ID(ctx)
		if err != nil {
			return err
		}
		renovate = renovate.WithSecretVariable(ele, client.Secret(secret))
	}

	// write dditional env variables
	for key, val := range opts.Env {
		renovate = renovate.WithEnvVariable(key, val)
	}

	renovate = renovate.WithEnvVariable("RENOVATE_PLATFORM", opts.Platform)
	renovate = renovate.WithEnvVariable("RENOVATE_EXTENDS", "github>whitesource/merge-confidence:beta")
	renovate = renovate.WithEnvVariable("RENOVATE_REQUIRE_CONFIG", "true")
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

	_, err = renovate.Exec().Stdout(ctx)
	if err != nil {
		panic(err)
	}

	return nil
}

func createImageString(img image) string {
	return fmt.Sprintf("%s:%s", img.Name, img.Version)
}
