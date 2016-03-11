package oxford

import (
	"fmt"
	"net/http"
	"strconv"

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

type photoURLType struct {
	URL string `json:"url"`
}

type photoLocalType struct {
	URL string `json:"url"`
}

type faceResponseType struct {
	PersistedFaceID string `json:"persistedFaceId"`
}

func (f face) AddFace(faceListID string, imageFileName string) (persistedFaceID string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID + "/persistedFaces"
	imageByteArray, err := imageToByteArray(imageFileName)
	fmt.Println("------------------")
	//fmt.Println(imageByteArray)
	fmt.Println(byteArrayToBase64(imageByteArray))
	fmt.Println("------------------")
	photo := photoLocalType{URL: byteArrayToBase64(imageByteArray)}
	if err != nil {
		return "", err
	}
	client, req := getPOSTClient(url, f.apiKey, photo, "application/octet-stream")

	fmt.Printf("AddFace Len: %d----->", len(photo.URL))

	req.Header.Add("Content-Length", strconv.Itoa(len(photo.URL)))
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	var faceResponse faceResponseType
	switch resp.StatusCode {
	case http.StatusOK:
		json.NewDecoder(resp.Body).Decode(&faceResponse)
		goops.Info(goops.Context(goops.C{"op": "AddFace", "result": "OK"}), "%s", resp.Status)
	default:
		var faceErrorResponse APIErrorResponse
		json.NewDecoder(resp.Body).Decode(&faceErrorResponse)
		goops.Info(goops.Context(goops.C{"op": "AddFace", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
		fmt.Print(toJSON(faceErrorResponse, pretty))
	}

	if err != nil {
		fmt.Println(err)
	}

	return toJSON(faceResponse, pretty), err

}

func (f face) AddFaceURL(faceListID string, photoURL string) (list string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID + "/persistedFaces"
	photo := photoURLType{URL: photoURL}
	client, req := getPOSTClient(url, f.apiKey, photo, "application/json")

	resp, err := client.Do(req)

	fmt.Printf("AddPhoto----->")

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "AddPhoto", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "AddPhoto", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	}

	if err != nil {
		fmt.Println(err)
	}

	var faceResponse faceResponseType
	switch resp.StatusCode {
	case http.StatusOK:
		json.NewDecoder(resp.Body).Decode(&faceResponse)
		goops.Info(goops.Context(goops.C{"op": "AddFace", "result": "OK"}), "%s", resp.Status)
	default:
		var faceErrorResponse APIErrorResponse
		json.NewDecoder(resp.Body).Decode(&faceErrorResponse)
		goops.Info(goops.Context(goops.C{"op": "AddFace", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
		fmt.Print(toJSON(faceErrorResponse, pretty))
	}

	if err != nil {
		fmt.Println(err)
	}

	return toJSON(faceResponse, pretty), err

}

func (f face) CreateFaceList(faceListID string) (id string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID
	fl := faceList{Name: faceListID, UserData: "Face List for Equinox"}
	client, req := getPUTClient(url, f.apiKey, fl, "application/json")

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

func (f face) GetFaceList() (list string, err error) {
	url := GetResource(Face, V1, "facelists")
	client, req := getGETClient(url, f.apiKey)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	}

	if err != nil {
		fmt.Println(err)
	}

	var listOfFacesList []faceList
	json.NewDecoder(resp.Body).Decode(&listOfFacesList)
	return toJSON(listOfFacesList, pretty), err

}

// NewFace creates a face client
func NewFace(key string) face {

	f := face{}
	f.apiKey = key
	return f
}
