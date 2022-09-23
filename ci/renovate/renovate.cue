package renovate

import (
	"dagger.io/dagger"
	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Run: {
	accessToken:        dagger.#Secret
	githubToken:        dagger.#Secret
	autodiscover:       *"false" | string
	autodiscoverFilter: *"" | string
	platform:           *"gitlab" | string
	repositories:       *"" | string
	version:            *"latest" | string
	logLevel:           *"error" | string

	_image: docker.#Pull & {
		source:      "renovate/renovate:\(version)"
		resolveMode: "preferLocal"
	}

	docker.#Run & {
		input:  _image.output
		always: true
		env: {
			RENOVATE_TOKEN:                accessToken
			GITHUB_COM_TOKEN:              githubToken
			RENOVATE_PLATFORM:             platform
			RENOVATE_EXTENDS:              "github>whitesource/merge-confidence:beta"
			RENOVATE_REQUIRE_CONFIG:       "true"
			RENOVATE_GIT_AUTHOR:           "Renovate Bot <bot@renovateapp.com>"
			RENOVATE_PIN_DIGEST:           "true"
			RENOVATE_DEPENDENCY_DASHBOARD: "false"
			RENOVATE_LABELS:               "renovate"
			RENOVATE_AUTODISCOVER:         autodiscover
			RENOVATE_AUTODISCOVER_FILTER:  autodiscoverFilter
			RENOVATE_REPOSITORIES:         repositories
			LOG_LEVEL:                     logLevel
		}
	}
}
