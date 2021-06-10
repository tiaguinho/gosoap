package gosoap

import (
	"encoding/xml"
	"testing"
)

var (
	mapParamsTests = []struct {
		Params Params
		Err    string
	}{
		{
			Params: Params{"": ""},
			Err:    "error expected: xml: start tag with no name",
		},
	}

	arrayParamsTests = []struct {
		Params ArrayParams
		Err    string
	}{
		{
			Params: ArrayParams{{"", ""}},
			Err:    "error expected: xml: start tag with no name",
		},
	}

	sliceParamsTests = []struct {
		Params SliceParams
		Err    string
	}{
		{
			Params: SliceParams{xml.StartElement{}, xml.EndElement{}},
			Err:    "error expected: xml: start tag with no name",
		},
	}
)

func TestClient_MarshalXML(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range mapParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}

func TestClient_MarshalXML2(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range arrayParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}

func TestClient_MarshalXML3(t *testing.T) {
	soap, err := SoapClient("https://kasapi.kasserver.com/soap/wsdl/KasAuth.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range mapParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}

func TestClient_MarshalXML4(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range sliceParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}

func TestSetCustomEnvelope(t *testing.T) {
	SetCustomEnvelope("soapenv", map[string]string{
		"xmlns:soapenv": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:tem": "http://tempuri.org/",
	})

	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	for _, test := range arrayParamsTests {
		_, err = soap.Call("checkVat", test.Params)
		if err == nil {
			t.Errorf(test.Err)
		}
	}
}
