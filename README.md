# dagger-ci-modules


## modules

### goreleaser

a module containing goreleaser.

``` go
import "github.com/tubenhirn/dagger-ci-modules/v2/goreleaser"

options := goreleaser.GoReleaserOpts{
	Source:     dir,
	Snapshot:   true,
	RemoveDist: true,
	Env: map[string]string{
	    "APP_VERSION": string(version),
	},
}

goreleaser.Release(context.Background(), options)
```

### semantic-release

### renovate

### golangci

## release

``` shell
dagger-cue do release
```
