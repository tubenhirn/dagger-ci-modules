package cimodules

import "fmt"

func createImageString(img image) string {
	return fmt.Sprintf("%s:%s", img.Name, img.Version)
}