package oxford

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/lentregu/Equinox/goops"
)

type smsType struct {
}

curl -X POST -d '{"to": ["tel:+34699218702"], 
"message": "Tu PIN es 8765", "from": "tel:22949;phone-context=+34"}' 
--header "Content-Type:application/json" http://81.45.59.59:8000/sms/v2/smsoutbound

type smsRequestType struct {
	from                     string `json:"from"`
	To                     []string `json:"to"`
	Message                 string `json:"message"`
}



func (f smsType) AddFaceURL(faceListID string, photoURL string) (list string, err error) {
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
// NewFace creates a face client
func NewSMS() smsType {

	sms := smsType{}
	return sms
}
