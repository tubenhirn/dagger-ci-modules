package goreleaser

import (
	"dagger.io/dagger"

	"universe.dagger.io/go"
)

// Like #ReleaseBase, but with a pre-configured container image.
#Release: #ReleaseBase & {
	_image: #Image
	image:  _image.output
}

// Release Go binaries using GoReleaser
#ReleaseBase: {
	// Source code
	source: dagger.#FS

	// Don't publish or announce the release
	dryRun: bool | *false

	// Build a snapshot instead of a tag
	snapshot: bool | *false

	// Remove dist dir
	removeDist: bool | *false

	// env vars
	env: [string]: string | dagger.#Secret

	go.#Container & {
		name:     "goreleaser"
		"source": source

		entrypoint: [] // Support images that does not set goreleaser as the entrypoint
		env: env
		command: {
			name: "goreleaser"

			flags: {
				if dryRun {
					"--skip-publish":  true
					"--skip-announce": true
				}

				if snapshot {
					"--snapshot": true
				}

				if removeDist {
					"--rm-dist": true
				}
			}
		}
		export: directories: "/src/dist": _
	}
}
