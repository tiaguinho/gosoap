package gosoap_test

import (
	"encoding/xml"
	"github.com/tiaguinho/gosoap"
	"reflect"
	"testing"
)

//
func TestSoapClient(t *testing.T) {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	if reflect.DeepEqual(soap, &gosoap.Client{}) {
		t.Errorf("got: \n %s", soap)
	}
}

//
type GetGeoIPResponse struct {
	GetGeoIPResult GetGeoIPResult
}

//
type GetGeoIPResult struct {
	ReturnCode        string
	IP                string
	ReturnCodeDetails string
	CountryName       string
	CountryCode       string
}

//
func TestClient_Call(t *testing.T) {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	b, err := soap.Call("GetGeoIP", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	var r GetGeoIPResponse
	xml.Unmarshal(b, &r)

	if r.GetGeoIPResult.CountryCode != "USA" {
		t.Errorf("error: %+v", r)
	}
}

//
type GetWhoISResponse struct {
	GetWhoISResult string
}

//
func TestClient_Call2(t *testing.T) {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/whois.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"HostName": "github.com",
	}

	b, err := soap.Call("GetWhoIS", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	var r GetWhoISResponse
	xml.Unmarshal(b, &r)

	if r.GetWhoISResult == "" {
		t.Errorf("error: %+v", r)
	}
}
