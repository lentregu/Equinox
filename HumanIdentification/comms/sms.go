package comms

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/lentregu/Equinox/goops"
)

type smsType struct {
}

// curl -X POST -d '{"to": ["tel:+34699218702"],
// "message": "Tu PIN es 8765", "from": "tel:22949;phone-context=+34"}'
// --header "Content-Type:application/json" http://81.45.59.59:8000/sms/v2/smsoutbound

type smsRequestType struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Message string   `json:"message"`
}

type smsResponseType struct {
	ID string `json:"id"`
}

// SendSMS is ...
func (f smsType) SendSMS(text string) (string, error) {

	url := "https://dev.mobileconnect.pdi.tid.es/es/sms/v2/smsoutbound"
	dst := []string{"tel:+34699218702"}

	smsBody := smsRequestType{From: "tel:22949;phone-context=+34", To: dst, Message: text}
	client, req := getClient(url, smsBody, "application/json")

	resp, err := client.Do(req)

	fmt.Printf("Sending SMS----->")

	if err != nil {
		return "", err
	}

	var smsResponse smsResponseType
	switch resp.StatusCode {
	case http.StatusOK:
		json.NewDecoder(resp.Body).Decode(&smsResponse)
		goops.Info("SMS ID: %s", smsResponse.ID)
		goops.Info(goops.Context(goops.C{"op": "SendSMS", "result": "OK"}), "%s", resp.Status)
	default:
		goops.Info(goops.Context(goops.C{"op": "SendSMS", "result": "NOK"}))
	}

	if err != nil {
		fmt.Println(err)
	}

	return smsResponse.ID, err

}

func getClient(url string, body interface{}, contentType string) (*http.Client, *http.Request) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	bodyJSON, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBufferString(fmt.Sprintf("%s", bodyJSON)))
	req.Header.Add("Content-Type", contentType)

	return client, req
}

// NewSMS creates a face client
func NewSMS() smsType {

	sms := smsType{}
	return sms
}
