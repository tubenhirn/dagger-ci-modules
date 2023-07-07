package cimodules

import "fmt"

func createImageString(defaultImg Image, customImage Image) string {
	image := defaultImg

	if customImage.Suffix != "" {
		image.Suffix = fmt.Sprintf("-%s", customImage.Suffix)
	}
	if customImage.Name != "" {
		image.Name = customImage.Name
	}
	if customImage.Version != "" {
		image.Version = customImage.Version
	}
	return fmt.Sprintf("%s:%s%s", image.Name, image.Version, image.Suffix)
}
