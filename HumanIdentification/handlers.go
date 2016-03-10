package main

import (
	"Equinox/oxford"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Logger is a handler to wrap other handlers and log basic request parameters
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Info(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func detect(w http.ResponseWriter, r *http.Request) {

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
