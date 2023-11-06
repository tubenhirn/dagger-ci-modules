package cimodules

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var customImage = Image{
	Name:    "custom",
	Version: "1.2.3",
	Suffix:  "suffix",
}

func TestCreateImageString(t *testing.T) {
	assert.Equal(t, "renovate/renovate:37.42.0", createImageString(defaultRenovateImage, Image{}))
	assert.Equal(t, "custom:1.2.3-suffix", createImageString(defaultRenovateImage, customImage))
}
