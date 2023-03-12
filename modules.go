package dagger_ci_modules

import (
	"context"

	"dagger.io/dagger"
	"github.com/tubenhirn/dagger-ci-modules/goreleaser"
	"github.com/tubenhirn/dagger-ci-modules/renovate"
	"github.com/tubenhirn/dagger-ci-modules/semanticrelease"
)

func Release(ctx context.Context, client dagger.Client, opts goreleaser.GoReleaserOpts) {
	goreleaser.Release(ctx, client, opts)
}

func Renovate(ctx context.Context, client dagger.Client, opts renovate.RenovateOpts){
	renovate.Renovate(ctx, client, opts)
}

func SemanticRelease(ctx context.Context, client dagger.Client, opts semanticrelease.SemanticOpts) {
	semanticrelease.Semanticrelease(ctx, client, opts)
}
