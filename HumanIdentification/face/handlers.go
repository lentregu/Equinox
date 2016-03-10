package face

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	// PrimaryKey for FaceAPI
	PrimaryKey = "567c560aa85245418459b82634bc7a98"
	// SecondaryKey for FaceAPI
	SecondaryKey = "4c1a4e7a02104577b045a2d046b20d29"
)

// Index is the welcome handler
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Detect is a handler to detect faces
func Detect(w http.ResponseWriter, r *http.Request) {

	//detectReq()
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
