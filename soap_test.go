package gosoap

import (
	"fmt"
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
			URL: "http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
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

type CheckVatResponse struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	RequestDate string `xml:"requestDate"`
	Valid       string `xml:"valid"`
	name        string `xml:"name"`
	Address     string `xml:"address"`
}

var (
	r CheckVatResponse

	params = Params{}
)

func TestClient_Call(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl")
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	params["vatNumber"] = "6388047V"
	params["countryCode"] = "IE"
	err = soap.Call("", params)
	if err == nil {
		t.Errorf("method is empty")
	}

	err = soap.Unmarshal(&r)
	if err == nil {
		t.Errorf("body is empty")
	}

	err = soap.Call("checkVat", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	fmt.Println(string(soap.GetLastRequest()))
	fmt.Println(string(soap.Body))
	soap.Unmarshal(&r)
	if r.CountryCode != "IE" {
		t.Errorf("error: %+v", r)
	}

	c := &Client{}
	err = c.Call("", Params{})
	if err == nil {
		t.Errorf("error expected but nothing got.")
	}

	c.WSDL = "://test."

	err = c.Call("checkVat", params)
	if err == nil {
		t.Errorf("invalid WSDL")
	}
}

func TestClient_doRequest(t *testing.T) {
	c := &Client{}

	_, err := c.doRequest("")
	if err == nil {
		t.Errorf("body is empty")
	}

	_, err = c.doRequest("://teste.")
	if err == nil {
		t.Errorf("invalid WSDL")
	}
}
