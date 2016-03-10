package oxford

import (
	"path/filepath"

	"github.com/lentregu/Equinox/goops"
)

type face struct {
}

func (f face) detect(img string) bool {
	url := GetResource(Face, V1, "detect")
	goops.Info("url: %s", url)
	return true
}

func (f face) verify(img string) bool {
	url := GetResource(Face, V1, "detect")
	goops.Info("url: %s", url)
	return true
}

func (f face) createFaceList(faceListID string) (id string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = filepath.Join(url, faceListID)
	goops.Info("url: %s", url)
	return "id", nil
}

// New creates a face client
func New() face {
	return face{}
}
