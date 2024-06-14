package cimodules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFlags(t *testing.T) {
	assert.Equal(t, []string(nil), createFlags(GoReleaserOpts{}))
	assert.Equal(t, []string{"--snapshot", "--clean"}, createFlags(GoReleaserOpts{Snapshot: true, RemoveDist: true}))
}
