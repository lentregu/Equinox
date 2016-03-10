package oxford

import (
	"fmt"
	"net/url"
	"path/filepath"
)

// APIType is a type for the different apis
type APIType int

const (
	// Face represents the face api
	Face APIType = iota
	// SpeakerRecognition represents the SpeakerRecognition api
	SpeakerRecognition
)

const (
	apiURL string = "https://api.projectoxford.ai"
	// V1 is the v1.0 version
	V1 string = "v1.0"
)

var apis map[APIType]string

func init() {

	apis = map[APIType]string{
		Face:               "face",
		SpeakerRecognition: "spid",
	}

}

// GetResource builds a resource
func GetResource(apiType APIType, version string, resource string) string {
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = filepath.Join(apis[apiType], version, resource)
	urlStr := fmt.Sprintf("%v", u)
	return urlStr
}
