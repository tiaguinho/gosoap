package gosoap

import (
	"testing"
)

var (
	tests = []struct {
		Params Params
		Err    string
	}{
		{
			Params: Params{},
			Err:    "params size is empty",
		},
		{
			Params: Params{"": ""},
			Err:    "error expected: xml: start tag with no name",
		},
	}
)

func TestClient_MarshalXML(t *testing.T) {
	soap, err := SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range tests {
		err = soap.Call("GetGeoIP", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}
