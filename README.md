# dagger-ci-modules


## modules

### goreleaser

a module providing goreleaser. https://github.com/goreleaser/goreleaser

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

a module providing renovate. https://github.com/renovatebot/renovate

``` go
import "github.com/tubenhirn/dagger-ci-modules/v2/renovate"

options := renovate.RenovateOpts{
    Platform: "github",
    Autodiscover: false,
    AutodiscoverFilter: "",
    Repositories: "tubenhirn/dagger-ci-modules",
    Env: map[string]string{},
    Secret: []string{
        "RENOVATE_TOKEN", "GITHUB_TOKEN"
    },
    LogLevel: "debug",
}

renovate.Renovate(context.Background(), options)
```

### golangci

## release

``` shell
dagger-cue do release
```
