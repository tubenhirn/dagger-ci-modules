package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/tubenhirn/dagger-ci-modules/v5"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}

	defer client.Close()

	platform := flag.String("platform", "", "the string of the platform to run on.")
	flag.Parse()

	fmt.Println("running with flags:", "\nplatform", *platform)

	var secrets = make(map[string]dagger.SecretID)
	switch *platform {
	case "github":
		token := os.Getenv("GITHUB_TOKEN")
		githubTokenId, err := client.SetSecret("GITHUB_TOKEN", token).ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId
	case "gitlab":
		token := os.Getenv("GITLAB_TOKEN")
		gitlabTokenId, err := client.SetSecret("GITHUB_TOKEN", token).ID(ctx)
		client.Container().EnvVariable(ctx, "GITLAB_TOKEN")
		if err != nil {
			panic(err)
		}
		secrets["GITLAB_TOKEN"] = gitlabTokenId
	case "git":
		token := os.Getenv("GITHUB_TOKEN")
		githubTokenId, err := client.SetSecret("GITHUB_TOKEN", token).ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId
	default:
		error := errors.New("flag platform missing.")
		panic(error)
	}

	dir, _ := os.Getwd()

	tmpReleaseImage := cimodules.Image{
		Name:    "tubenhirn/semantic-release-github",
		Version: "v4.1.7",
	}
	options := cimodules.SemanticOpts{
		Source:   dir,
		Platform: *platform,
		Env:      map[string]string{},
		Secret:   secrets,
		Image:    tmpReleaseImage,
	}

	if err := cimodules.Semanticrelease(ctx, *client, options); err != nil {
		fmt.Println(err)
	}
}
