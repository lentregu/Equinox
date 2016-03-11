package oxford

import (
	"bytes"
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

type faceType struct {
	PersistedFaceID string `json:"persistedFaceId"`
	UserData        string `json:"userData"`
}
type faceListContent struct {
	FaceListID     string `json:"faceListId"`
	Name           string `json:"name"`
	UserData       string `json:"userData"`
	persistedFaces []faceType
}

func (f face) Verify(img string) bool {
	url := GetResource(Face, V1, "detect")
	goops.Info("url: %s", url)
	return true
}

// PhotoURLType is..
type PhotoURLType struct {
	URL string `json:"url"`
}

// PhotoLocalType  is ...
type PhotoLocalType struct {
	URL string `json:"url"`
}

type faceResponseType struct {
	PersistedFaceID string `json:"persistedFaceId"`
}

type faceSimilarRequestType struct {
	FaceID                     string `json:"faceId"`
	FaceListID                 string `json:"faceListId"`
	MaxNumOfCandidatesReturned int    `json:"maxNumOfCandidatesReturned"`
}

type faceSimilarResponseType struct {
	PersistedFaceID string  `json:"persistedFaceId"`
	Confidence      float64 `json:"confidence"`
}

type faceDetectInfo struct {
	FaceID string `json:"faceId"`
	//El resto no me interesan
}

func (f face) Detect(photoURL string) (string, error) {
	url := GetResource(Face, V1, "detect")
	photo := PhotoURLType{URL: photoURL}
	client, req := getPOSTClient(url, f.apiKey, photo, "application/json")
	req.URL.Query().Add("returnFaceId", "true")
	resp, err := client.Do(req)

	var faceID string

	if err != nil {
		return faceID, err
	}

	var faceDetectResponse []faceDetectInfo

	switch resp.StatusCode {
	case http.StatusOK:
		json.NewDecoder(resp.Body).Decode(&faceDetectResponse)
		goops.Info(goops.Context(goops.C{"op": "Detect", "result": "OK"}), "%s", resp.Status)
		faceID = faceDetectResponse[0].FaceID
	default:
		var faceErrorResponse APIErrorResponse
		json.NewDecoder(resp.Body).Decode(&faceErrorResponse)
		goops.Info(goops.Context(goops.C{"op": "Detect", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
		fmt.Print(toJSON(faceErrorResponse, pretty))
	}

	return faceID, nil

}

func (f face) FindSimilar(faceID string, faceListID string) (bool, error) {
	url := GetResource(Face, V1, "findsimilars")
	faceSimilarBody := faceSimilarRequestType{FaceID: faceID, FaceListID: faceListID, MaxNumOfCandidatesReturned: 5}
	client, req := getPOSTClient(url, f.apiKey, faceSimilarBody, "application/json")

	var similarList []faceSimilarResponseType
	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		json.NewDecoder(resp.Body).Decode(&similarList)
		fmt.Print(toJSON(similarList, pretty))
		goops.Info(goops.Context(goops.C{"op": "FindSimilar", "result": "OK"}), "%s", resp.Status)
	default:
		var similarErrorResponse APIErrorResponse
		json.NewDecoder(resp.Body).Decode(&similarErrorResponse)
		goops.Info(goops.Context(goops.C{"op": "FindSimilar", "result": "NOK"}))
		goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
		fmt.Print(toJSON(similarErrorResponse, pretty))
	}

	if err != nil {
		fmt.Println(err)
	}

	found := false
	for _, similar := range similarList {
		if similar.Confidence > 0.5 {
			found = true
		}
	}

	return found, err
}

func (f face) AddFace(faceListID string, imageFileName string) (persistedFaceID string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID + "/persistedFaces"
	imageByteArray, err := imageToByteArray(imageFileName)
	fmt.Println("------------------")
	//fmt.Println(imageByteArray)
	fmt.Println(byteArrayToBase64(imageByteArray))
	fmt.Println("------------------")
	photo := PhotoLocalType{URL: byteArrayToBase64(imageByteArray)}
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
	photo := PhotoURLType{URL: photoURL}
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

func (f face) GetFacesInAList(faceListID string) (list string, err error) {
	url := GetResource(Face, V1, "facelists")
	url = url + "/" + faceListID
	client, req := getGETClient(url, f.apiKey)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "NOK"}))
	}

	if err != nil {
		fmt.Println(err)
	}

	facesInAList := faceListContent{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Printf("--->%s", buf)
	json.NewDecoder(resp.Body).Decode(&facesInAList)
	return toJSON(facesInAList, pretty), err

}

// NewFace creates a face client
func NewFace(key string) face {

	f := face{}
	f.apiKey = key
	return f
}
