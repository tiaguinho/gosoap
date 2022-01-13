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
	CountryCode string   `xml:"countryCode,omitempty"`
	VatNumber   string   `xml:"vatNumber"`
	TraderName  string   `xml:"traderName,omitempty"`
}
type checkVatApproxResponse struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	Valid       bool   `xml:"valid"`
	TraderName  string `xml:"traderName,omitempty"`
}

var encoderParamsTests = []struct {
	Desc     string
	WSDL     string
	Params   *checkVatApprox
	Response *checkVatApproxResponse
	Err      string
}{
	{
		Desc:   "Fetch a non-existent VAT number",
		WSDL:   "http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
		Params: &checkVatApprox{CountryCode: "fr", VatNumber: "invalid"},
		Response: &checkVatApproxResponse{
			CountryCode: "FR", VatNumber: "invalid", Valid: false, TraderName: "---",
		},
	},
	{
		Desc:   "Fetch a valid VAT number",
		WSDL:   "http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
		Params: &checkVatApprox{CountryCode: "fr", VatNumber: "45327920054"},
		Response: &checkVatApproxResponse{
			CountryCode: "FR", VatNumber: "45327920054", Valid: true, TraderName: "SAS EUROMEDIA",
		},
	},
	{
		Desc:   "Fetch with empty params",
		WSDL:   "http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
		Params: &checkVatApprox{},
		Err:    `[soap:Server]: Invalid_input | Detail: `,
	},
}

func (cva *checkVatApprox) SoapBuildRequest() *Request {
	r := NewRequest("checkVatApprox", cva)
	r.UseXMLEncoder = true
	return r
}
func TestClient_MarshalWithEncoder(t *testing.T) {
	for _, test := range encoderParamsTests {
		soap, err := SoapClient(test.WSDL, nil)
		if err != nil {
			t.Errorf("%s: error not expected creating client: %s", test.Desc, err)
			continue
		}

		resp, err := soap.CallByStruct(test.Params)
		if err != nil {
			t.Errorf("%s: error not expected calling API: %s", test.Desc, err)
			continue
		}

		var actualResponse checkVatApproxResponse
		err = resp.Unmarshal(&actualResponse)
		if test.Err != "" {
			if err == nil {
				t.Errorf("%s: expected error, but got response: %#v", test.Desc, actualResponse)
				continue
			} else if err.Error() != test.Err {
				t.Errorf("%s: error doesn't match expectation: %s", test.Desc, err)
			}
		} else {
			if err != nil {
				t.Errorf("%s: unmarshal error not expected: %s", test.Desc, err)
				continue
			} else if actualResponse != *test.Response {
				t.Errorf("%s: response doesn't match expectation: %#v", test.Desc, actualResponse)
				continue
			}

		}
	}
}
