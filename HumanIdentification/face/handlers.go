package face

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lentregu/Equinox/oxford"
)

// Index is the welcome handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Detect is a handler to detect faces
func Detect(w http.ResponseWriter, r *http.Request) {

	detectReq()
	info := InfoFaceDetection{
		Name:      "Gonzalo",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(info); err != nil {
		panic(err)
	}
}

func detectReq() {

	resource := oxford.GetResource(oxford.Face, oxford.V1, "detect")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	//_, err := client.Get("https://https://api.projectoxford.ai/face/v1.0/detect")
	_, err := client.Get(resource)
	if err != nil {
		fmt.Println(err)
	}

}
