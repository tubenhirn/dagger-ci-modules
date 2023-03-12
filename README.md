# dagger-ci-modules

## modules

### goreleaser

a module providing goreleaser. https://github.com/goreleaser/goreleaser

```go
import (
    "dagger.io/dagger"
    "github.com/tubenhirn/dagger-ci-modules/v4"
)

// a context
ctx := context.Background()

// initialize Dagger client
client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
if err != nil {
    panic(err)
}

defer client.Close()

dir, _ := os.Getwd()

options := cimodules.GoReleaserOpts{
    Source:     dir,
    Snapshot:   true,
    RemoveDist: true,
    Env: map[string]string{
        "APP_VERSION": string(version),
    },
}

cimodules.Release(ctx, *client, options)
```

### semantic-release

a module providing semantic-release. https://github.com/semantic-release/github

```go
import (
    "dagger.io/dagger"
    "github.com/tubenhirn/dagger-ci-modules/v4"
)

// a context
ctx := context.Background()

// initialize Dagger client
client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
if err != nil {
    panic(err)
}

defer client.Close()

var secrets = make(map[string]dagger.SecretID)
githubTokenId, err = client.Host().EnvVariable("GITHUB_ACCESS_TOKEN").Secret().ID(ctx)
if err != nil {
    panic(err)
}
secrets["GITHUB_TOKEN"] = githubTokenId

dir, _ := os.Getwd()

options := cimodules.SemanticOpts{
    Source:   dir,
    Platform: "github",
    Env:      map[string]string{},
    Secret:   secrets,
}

if err := cimodules.Semanticrelease(ctx, *client, options); err != nil {
    fmt.Println(err)
}
```

### renovate

a module providing renovate. https://github.com/renovatebot/renovate

```go
import (
    "dagger.io/dagger"
    "github.com/tubenhirn/dagger-ci-modules/v4"
)

// a context
ctx := context.Background()

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

options := cimodules.RenovateOpts{
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

cimodules.Renovate(ctx, *client, options)
```

## create a new release

use the included `Makefile` to run the release job.

```shell
make release
```
