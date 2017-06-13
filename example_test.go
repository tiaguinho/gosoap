package gosoap_test

import (
	"github.com/tiaguinho/gosoap"
	"testing"
)

type GetGeoIPResponse struct {
	GetGeoIPResult GetGeoIPResult
}

type GetGeoIPResult struct {
	ReturnCode        string
	IP                string
	ReturnCodeDetails string
	CountryName       string
	CountryCode       string
}

var (
	r GetGeoIPResponse
)

func TestClient_Call(t *testing.T) {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	err = soap.Call("GetGeoIP", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	soap.Unmarshal(&r)
	if r.GetGeoIPResult.CountryCode != "USA" {
		t.Errorf("error: %+v", r)
	}
}
