package gosoap

import (
	"encoding/xml"
	"fmt"
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
		"xmlns:tem":     "http://tempuri.org/",
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

type checkVatApprox struct {
	XMLName     xml.Name `xml:"urn:ec.europa.eu:taxud:vies:services:checkVat:types checkVatApprox"`
	CountryCode string   `xml:"countryCode"`
	VatNumber   string   `xml:"vatNumber"`
	TraderName  string   `xml:"traderName,omitempty"`
}
type checkVatApproxResponse struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	Valid       bool   `xml:"valid"`
	TraderName  string `xml:"traderName,omitempty"`
}

func (cva *checkVatApprox) SoapBuildRequest() *Request {
	r := NewRequest("checkVatApprox", cva)
	// if err!=nil{
	// 	t.Errorf("error not expected: %s", err)
	// }
	r.UseXMLEncoder = true
	return r
}
func TestClient_MarshalWithEncoder(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	resp, err := soap.CallByStruct(&checkVatApprox{CountryCode: "fr", VatNumber: "67586586"})
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	var cvaResp checkVatApproxResponse
	err = resp.Unmarshal(&cvaResp)
	if err != nil {
		t.Errorf("unmarshal error not expected: %s", err)
	}

	fmt.Printf("\n  resp: %#v\n", resp)
	fmt.Printf("  payload: %s\n", resp.Payload)
	fmt.Printf("  body: %s\n", resp.Body)
	fmt.Printf("  err: %s\n", err)
	fmt.Printf("   unmarshaled: %#v\n", cvaResp)
	expectResp:=checkVatApproxResponse{CountryCode:"FR",VatNumber:"67586586",Valid:false,TraderName:"---"}
	if cvaResp!= expectResp{
		t.Errorf("got unexpected response: %#v",cvaResp)
	}
}
