package cimodules

import "fmt"

func createImageString(img image) string {
	suffix := ""
	if img.Suffix != "" {
		suffix = fmt.Sprintf("-%s", img.Suffix)
	}
	return fmt.Sprintf("%s:%s%s", img.Name, img.Version, suffix)
}
