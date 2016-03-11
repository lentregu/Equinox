package face

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lentregu/Equinox/HumanIdentification/comms"
	"github.com/lentregu/Equinox/oxford"
)

const (
	// PrimaryKey for FaceAPI
	PrimaryKey = "567c560aa85245418459b82634bc7a98"
	// SecondaryKey for FaceAPI
	SecondaryKey = "4c1a4e7a02104577b045a2d046b20d29"
)

type findSimilarRequestType struct {
	URL        string `json:"url"`
	FaceListID string `json:"faceListID"`
}

// Index is the welcome handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// FindSimilar is a handler to detect faces
func FindSimilar(w http.ResponseWriter, r *http.Request) {
	requestBody := findSimilarRequestType{}

	json.NewDecoder(r.Body).Decode(&requestBody)

	faceService := oxford.NewFace("567c560aa85245418459b82634bc7a98")
	faceID, _ := faceService.Detect(requestBody.URL)

	fmt.Printf("El faceID es: %s\n", faceID)

	if isSimilar, err := faceService.FindSimilar(faceID, requestBody.FaceListID); err != nil {
		fmt.Printf("Error %v", err)
	} else {

		if isSimilar {
			sms := comms.NewSMS(comms.Smppadapter)
			sms.SendSMS("PERSONA AUTORIZADA")
			fmt.Println("PERSONA AUTORIZADA")
		} else {
			fmt.Println("PERSONA NO AUTORIZADA")
		}
	}

}
