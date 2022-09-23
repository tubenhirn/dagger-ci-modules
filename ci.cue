package ci

import (
	"dagger.io/dagger"

	"github.com/tubenhirn/dagger-ci-modules/releasing"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents: dagger.#FS

	client: env: {
		GITLAB_TOKEN: dagger.#Secret
		GITHUB_TOKEN: dagger.#Secret
	}

	actions: {
		_source: client.filesystem["."].read.contents

		release: {
			semanticRelease: releasing.#Release & {
				sourcecode: _source
				authToken:  client.env.GITLAB_TOKEN
				version:    "v2.5.0"
			}
		}
	}
}
