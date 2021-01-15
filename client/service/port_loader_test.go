package service

import (
	"ports/client/domain"
	"strings"
	"testing"
)

const json = `{
"AEAJM": {
"name": "Ajman",
"city": "Ajman",
"country": "United Arab Emirates",
"alias": [],
"regions": [],
"coordinates": [
55.5136433,
25.4052165
],
"province": "Ajman",
"timezone": "Asia/Dubai",
"unlocs": [
"AEAJM"
],
"code": "52000"
},
"AEAUH": {
"name": "Abu Dhabi",
"coordinates": [
54.37,
24.47
],
"city": "Abu Dhabi",
"province": "Abu Z¸aby [Abu Dhabi]",
"country": "United Arab Emirates",
"alias": [],
"regions": [],
"timezone": "Asia/Dubai",
"unlocs": [
"AEAUH"
],
"code": "52001"
}
}
`

type MockPortSenderService struct {
	T *testing.T
}
func ( s MockPortSenderService) SendPort(port domain.Port) error{

	//Assert that expected port is recieved othersise fail the test.
	//if port = json
	return nil
}


func TestLoader(t *testing.T) {
	mockService := MockPortSenderService{t}
	s := NewGRPCPortLoaderService()
	s.LoadPorts(strings.NewReader(json),mockService)
}


/*TODO: Test Cases to implement
Test For decode json - Happy Path
Test for decoding json - Incorrect input - fail gracefully.

Test for sending - Happy Path
Teet for sending - Failure to send to server. Do we want to fail service  if we are unable to initialise, or can we
work with a partially loaded dataset?

 */