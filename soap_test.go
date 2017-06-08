package gosoap_test

import (
	"encoding/xml"
	"github.com/tiaguinho/gosoap"
	"reflect"
	"testing"
)

var (
	wsdl = "http://www.webservicex.net/geoipservice.asmx?wsdl"
	url  = "http://www.webservicex.net/"
)

//
func TestSoapClient(t *testing.T) {
	soap := gosoap.SoapClient(wsdl, url)

	if reflect.DeepEqual(soap, gosoap.Client{WSDL: wsdl, URL: url}) {
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
	soap := gosoap.SoapClient(wsdl, url)

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
