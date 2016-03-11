package face

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/lentregu/Equinox/goops"
	"github.com/lentregu/Equinox/oxford"
)

// APIError represents an error in oxford responses
type APIError struct {
	Err oxford.Error `json:"error"`
}

// FaceService is ...
type FaceService struct {
}

// CreateFaceList is ...
func (f FaceService) CreateFaceList(faceListID string) (string, error) {
	resource := oxford.GetResource(oxford.Face, oxford.V1, "facelists")
	resource = resource + "/" + faceListID

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, _ := http.NewRequest("PUT", resource, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", oxford.AzureSubscriptionID)

	//fmt.Print(req)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	//fmt.Print(resp)
	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "NOK"}), "%s", resp.Status)
	}

	if err != nil {
		fmt.Println(err)
	}

	return faceListID, err
}

// func detect() {

// 	resource := oxford.GetResource(oxford.Face, oxford.V1, "detect")

// 	r.Header.Add("Content-Type", "application/json")
// 	r.Header.Add("Ocp-Apim-Subscription-Key")

// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	client := &http.Client{Transport: tr}

// 	//_, err := client.Get("https://https://api.projectoxford.ai/face/v1.0/detect")
// 	_, err := client.Get(resource)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }
