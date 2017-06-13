package gosoap

import (
	"testing"
)

var (
	scts = []struct {
		URL string
		Err bool
	}{
		{
			URL: "://www.server",
			Err: false,
		},
		{
			URL: "",
			Err: false,
		},
		{
			URL: "http://www.webservicex.net/geoipservice.asmx?WSDL",
			Err: true,
		},
	}
)

func TestSoapClient(t *testing.T) {
	for _, sct := range scts {
		_, err := SoapClient(sct.URL)
		if err != nil && sct.Err {
			t.Errorf("URL: %s - error: %s", sct.URL, err)
		}
	}
}

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

	params = Params{}
)

func TestClient_Call(t *testing.T) {
	soap, err := SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	err = soap.Call("GetGeoIP", params)
	if err == nil {
		t.Errorf("params is empty")
	}

	params["IPAddress"] = "8.8.8.8"
	err = soap.Call("", params)
	if err == nil {
		t.Errorf("method is empty")
	}

	err = soap.Unmarshal(&r)
	if err == nil {
		t.Errorf("body is empty")
	}

	err = soap.Call("GetGeoIP", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	soap.Unmarshal(&r)
	if r.GetGeoIPResult.CountryCode != "USA" {
		t.Errorf("error: %+v", r)
	}

	c := &Client{}
	err = c.Call("", Params{})
	if err == nil {
		t.Errorf("error expected but nothing got.")
	}

	c.WSDL = "://test."

	err = c.Call("GetGeoIP", params)
	if err == nil {
		t.Errorf("invalid WSDL")
	}
}

func TestClient_doRequest(t *testing.T) {
	c := &Client{}

	_, err := c.doRequest()
	if err == nil {
		t.Errorf("body is empty")
	}

	c.WSDL = "://teste."
	_, err = c.doRequest()
	if err == nil {
		t.Errorf("invalid WSDL")
	}
}
