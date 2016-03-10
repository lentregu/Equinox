package oxford

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/lentregu/Equinox/goops"
)

type face struct {
	apiKey string
}

type faceList struct {
	FaceListID string `json:"faceListId, omitempty"`
	Name       string `json:"name"`
	UserData   string `json:"userData"`
}

func (f face) Detect(img string) bool {
	url := GetResource(Face, V1, "detect")
	goops.Info("url: %s", url)
	return true
}

func (f face) Verify(img string) bool {
	url := GetResource(Face, V1, "detect")
	goops.Info("url: %s", url)
	return true
}

func (f face) GetFaceList() (list string, err error) {
	url := GetResource(Face, V1, "facelists")
	client, req := getGETClient(url, f.apiKey)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	}

	if err != nil {
		fmt.Println(err)
	}

	var listOfFacesList []faceList
	json.NewDecoder(resp.Body).Decode(&listOfFacesList)
	return toJSON(listOfFacesList, pretty), err

}

func (f face) CreateFaceList(faceListID string) (id string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID
	fl := faceList{Name: faceListID, UserData: "Face List for Equinox"}
	client, req := getPUTClient(url, f.apiKey, fl)

	resp, err := client.Do(req)

	fmt.Println("-----------------")
	fmt.Print(req)
	fmt.Println()
	fmt.Println("-----------------")

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	}

	if err != nil {
		fmt.Println(err)
	}

	return faceListID, err
}

// NewFace creates a face client
func NewFace(key string) face {

	f := face{}
	f.apiKey = key
	return f
}
