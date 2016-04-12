package oxford

import (
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/lentregu/Equinox/goops"
)

type M map[string]string

type face struct {
	apiKey string
}

type faceList struct {
	FaceListID string `json:"faceListId,omitempty"`
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
	PersistedFaces []faceType
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

	resp, err := POST(url, M{"returnFaceId": "true"}, f.apiKey, nil, "application/json", photo)

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
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
		fmt.Print(toJSON(faceErrorResponse, pretty))
	}

	return faceID, nil

}

func (f face) FindSimilar(faceID string, faceListID string) (bool, error) {
	url := GetResource(Face, V1, "findsimilars")
	faceSimilarBody := faceSimilarRequestType{FaceID: faceID, FaceListID: faceListID, MaxNumOfCandidatesReturned: 5}

	var similarList []faceSimilarResponseType
	resp, err := POST(url, nil, f.apiKey, nil, "application/json", faceSimilarBody)

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
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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
	imageByteArray, err := fileToByteArray(imageFileName)

	if err != nil {
		return "", err
	}

	resp, err := POST(url, nil, f.apiKey, M{"Content-Length": strconv.Itoa(len(imageByteArray))}, "application/octet-stream", imageByteArray)

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
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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

	resp, err := POST(url, nil, f.apiKey, nil, "application/json", photo)

	fmt.Printf("AddPhoto----->")

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "AddPhoto", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "AddPhoto", "result": "NOK"}))
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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

	resp, err := PUT(url, nil, f.apiKey, nil, "application/json", fl)

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "createFaceList", "result": "NOK"}))
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
	}

	if err != nil {
		fmt.Println(err)
	}

	return faceListID, err
}

func (f face) GetFaceList() (list string, err error) {
	url := GetResource(Face, V1, "facelists")

	resp, err := GET(url, f.apiKey, nil, nil)

	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "GetFaceList", "result": "NOK"}))
		//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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

	resp, err := GET(url, f.apiKey, nil, nil)

	if err != nil {
		return "", err
	}

	//goops.Info("Status:%s|Request:%s", resp.Status, req.URL.RequestURI())
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
	json.NewDecoder(resp.Body).Decode(&facesInAList)
	return toJSON(facesInAList, pretty), err

}


// NewFace creates a face client
func NewFace(key string) face {

	f := face{}
	f.apiKey = key
	return f
}
