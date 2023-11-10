package cimodules

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var customImage = Image{
	Name:    "custom",
	Version: "1.2.3",
	Suffix:  "suffix",
}

func TestCreateImageString(t *testing.T) {
	//# renovate: datasource=docker depName=renovate/renovate versioning=docker
	renovateVersion := "37.53.1"
	assert.Equal(t, fmt.Sprintf("renovate/renovate:%s",renovateVersion), createImageString(defaultRenovateImage, Image{}))
	assert.Equal(t, "custom:1.2.3-suffix", createImageString(defaultRenovateImage, customImage))
	// test a failure case
	assert.NotEqual(t, "custom:1.2.3", createImageString(defaultRenovateImage, customImage))
}
