package releasing

import (
	"dagger.io/dagger"
	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Release: {
	authToken:  dagger.#Secret
	sourcecode: dagger.#FS
	version:    *"latest" | string

	_image: docker.#Pull & {
		source: "tubenhirn/semantic-release-gitlab:\(version)"
	}

	docker.#Run & {
		input: _image.output
		mounts: code: {
			dest:     "/src"
			contents: sourcecode
		}
		workdir: "/src"
		env: {
			GITLAB_TOKEN: authToken
		}
	}
}
