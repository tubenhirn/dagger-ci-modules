package releasing

import (
	"dagger.io/dagger"
	"universe.dagger.io/docker"
)

// create a release using semenatic-release
#Release: {
	authToken:  dagger.#Secret
	platform:   *"gitlab" | string
	sourcecode: dagger.#FS
	version:    *"latest" | string

	_image: docker.#Pull & {
		source: "tubenhirn/semantic-release-\(platform):\(version)"
	}

	docker.#Run & {
		input: _image.output
		mounts: code: {
			dest:     "/src"
			contents: sourcecode
		}
		workdir: "/src"
		env: {
			if platform == "gitlab" {
				GITLAB_TOKEN: authToken
			}
			if platform == "github" {
				GITHUB_TOKEN: authToken
			}
			if platform == "git" {
				GITHUB_TOKEN: authToken
			}
		}
	}
}
