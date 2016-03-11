package oxford

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	// AzureSubscriptionID is my azure subscription
	AzureSubscriptionID string = "70306775-8047-4d29-9540-679cc5412f0f"
)

// Error represents the structure of an oxford error
type oxfordError struct {
	StatusCode string `json:"code"`
	Message    string `json:"message"`
}

// APIErrorResponse is ...
type APIErrorResponse struct {
	Err oxfordError `json:"error"`
}

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

func parseError(body io.Reader) APIErrorResponse {
	err := APIErrorResponse{}
	json.NewDecoder(body).Decode(&err)
	return err
}

type printOption int

const (
	pretty printOption = iota
	normal
)

func toJSON(value interface{}, option printOption) string {

	var jsonValue []byte

	switch option {
	case pretty:
		jsonValue, _ = json.MarshalIndent(value, "", "\t")
	case normal:
		jsonValue, _ = json.Marshal(value)
	}

	return fmt.Sprintf("%s", jsonValue)
}

func getPOSTClient(url string, key string, body interface{}, contentType string) (*http.Client, *http.Request) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(toJSON(body, normal)))
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Ocp-Apim-Subscription-Key", key)

	return client, req
}

func getPUTClient(url string, key string, body interface{}, contentType string) (*http.Client, *http.Request) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("PUT", url, bytes.NewBufferString(toJSON(body, normal)))
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Ocp-Apim-Subscription-Key", key)

	return client, req
}

func getGETClient(url string, key string) (*http.Client, *http.Request) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Ocp-Apim-Subscription-Key", key)

	return client, req
}

func imageToByteArray(imageFileName string) ([]byte, error) {
	file, err := os.Open(imageFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes, err
}

func byteArrayToBase64(binaryByteArray []byte) string {
	imgBase64Str := base64.StdEncoding.EncodeToString(binaryByteArray)
	return imgBase64Str
}
