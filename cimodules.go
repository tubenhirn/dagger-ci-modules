package cimodules

import (
	"context"

	"dagger.io/dagger"
)

func Release(ctx context.Context, client dagger.Client, opts GoReleaserOpts) error{
	if err := release(ctx, client, opts); err != nil {
		return err
	}
	return nil
}

func Renovate(ctx context.Context, client dagger.Client, opts RenovateOpts) error {
	if err := renovate(ctx, client, opts); err != nil {
		return err
	}
	return nil
}

func Semanticrelease(ctx context.Context, client dagger.Client, opts SemanticOpts) error {
	if err := semanticrelease(ctx, client, opts); err != nil {
		return err
	}
	return nil
}
