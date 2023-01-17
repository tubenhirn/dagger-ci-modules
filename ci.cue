package ci

import (
	"dagger.io/dagger"

	"github.com/tubenhirn/dagger-ci-modules/releasing"
)

dagger.#Plan & {
	client: filesystem: ".": read: contents: dagger.#FS

	client: env: {
		GITHUB_TOKEN: dagger.#Secret
	}

	actions: {
		_source: client.filesystem["."].read.contents

		release: {
			semanticRelease: releasing.#Release & {
				sourcecode: _source
				platform:   "git"
				authToken:  client.env.GITHUB_TOKEN
				version:    "v2.9.0"
			}
		}
	}
}
