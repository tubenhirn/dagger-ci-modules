package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"dagger.io/dagger"
	"github.com/tubenhirn/dagger-ci-modules/semanticrelease"
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
		githubTokenId, err := client.Host().EnvVariable("GITHUB_TOKEN").Secret().ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId
	case "gitlab":
		gitlabTokenId, err := client.Host().EnvVariable("GITLAB_TOKEN").Secret().ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITLAB_TOKEN"] = gitlabTokenId
	case "git":
		githubTokenId, err := client.Host().EnvVariable("GITHUB_TOKEN").Secret().ID(ctx)
		if err != nil {
			panic(err)
		}
		secrets["GITHUB_TOKEN"] = githubTokenId
	default:
		error := errors.New("flag platform missing.")
		panic(error)
	}

	dir, _ := os.Getwd()

	options := semanticrelease.SemanticOpts{
		Source:   dir,
		Platform: *platform,
		Env:      map[string]string{},
		Secret:   secrets,
	}

	if err := semanticrelease.Semanticrelease(context.Background(), *client, options); err != nil {
		fmt.Println(err)
	}
}
