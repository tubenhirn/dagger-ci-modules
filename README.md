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
import (
    "dagger.io/dagger"
    "github.com/tubenhirn/dagger-ci-modules/v2/renovate"
)

// initialize Dagger client
client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
if err != nil {
    panic(err)
}

defer client.Close()

renovateTokenId, err = client.Host().EnvVariable("GITHUB_ACCESS_TOKEN").Secret().ID(ctx)
if err != nil {
    panic(err)
}

options := renovate.RenovateOpts{
    Platform: "github",
    Autodiscover: false,
    AutodiscoverFilter: "",
    Repositories: "tubenhirn/dagger-ci-modules",
    Env: map[string]string{},
    Secret: map[string]dagger.SecretID{
        "RENOVATE_TOKEN": renovateTokenId,
        "GITHUB_COM_TOKEN":   renovateTokenId,
    },
    LogLevel: "debug",
}

renovate.Renovate(context.Background(), client, options)
```

### golangci

## release

``` shell
dagger-cue do release
```
