package main

import "time"

// InfoFaceDetection contains the info return by the Detection faceAPI
type InfoFaceDetection struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
