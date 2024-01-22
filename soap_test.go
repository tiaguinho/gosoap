package gosoap

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

var (
	scts = []struct {
		URL    string
		Err    bool
		Client *http.Client
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
		{
			URL: "http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
			Err: true,
			Client: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			},
		},
	}
)

func TestSoapClient(t *testing.T) {
	for _, sct := range scts {
		_, err := SoapClient(sct.URL, nil)
		if err != nil && sct.Err {
			t.Errorf("URL: %s - error: %s", sct.URL, err)
		}
	}
}

func TestSoapClienWithClient(t *testing.T) {
	client, err := SoapClient(scts[3].URL, scts[3].Client)

	if client.HTTPClient != scts[3].Client {
		t.Errorf("HTTP client is not the same as in initialization: - error: %s", err)
	}

	if err != nil {
		t.Errorf("URL: %s - error: %s", scts[3].URL, err)
	}
}

type CheckVatRequest struct {
	CountryCode string
	VatNumber   string
}

func (r CheckVatRequest) SoapBuildRequest() *Request {
	return NewRequest("checkVat", Params{
		"countryCode": r.CountryCode,
		"vatNumber":   r.VatNumber,
	})
}

type CheckVatResponse struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	RequestDate string `xml:"requestDate"`
	Valid       string `xml:"valid"`
	Name        string `xml:"name"`
	Address     string `xml:"address"`
}

type CapitalCityResponse struct {
	CapitalCityResult string
}

type NumberToWordsResponse struct {
	NumberToWordsResult string
}

type WhoisResponse struct {
	WhoisResult string
}

var (
	rv CheckVatResponse
	rc CapitalCityResponse
	rn NumberToWordsResponse
	rw WhoisResponse

	params = Params{}
)

func TestClient_Call(t *testing.T) {
	soap, err := SoapClientWithConfig("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
		nil,
		&Config{Dump: true},
	)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	var res *Response

	params["vatNumber"] = "6388047V"
	params["countryCode"] = "IE"
	res, err = soap.Call("", params)
	if err == nil {
		t.Errorf("method is empty")
	}

	if res != nil {
		t.Errorf("body is empty")
	}

	res, err = soap.Call("checkVat", params)
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rv)
	if rv.CountryCode != "IE" {
		t.Errorf("error: %+v", rv)
	}

	soap, err = SoapClient("http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso?WSDL", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	res, err = soap.Call("CapitalCity", Params{"sCountryISOCode": "GB"})
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rc)

	if rc.CapitalCityResult != "London" {
		t.Errorf("error: %+v", rc)
	}

	soap, err = SoapClient("http://www.dataaccess.com/webservicesserver/numberconversion.wso?WSDL", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	res, err = soap.Call("NumberToWords", Params{"ubiNum": "23"})
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rn)

	if rn.NumberToWordsResult != "twenty three " {
		t.Errorf("error: %+v", rn)
	}

	soap, err = SoapClient("https://domains.livedns.co.il/API/DomainsAPI.asmx?WSDL", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	res, err = soap.Call("Whois", Params{"DomainName": "google.com"})
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rw)

	if rw.WhoisResult != "0" {
		t.Errorf("error: %+v", rw)
	}

	c := &Client{}
	_, err = c.Call("", Params{})
	if err == nil {
		t.Errorf("error expected but nothing got.")
	}

	c.SetWSDL("://test.")

	_, err = c.Call("checkVat", params)
	if err == nil {
		t.Errorf("invalid WSDL")
	}
}

type customLogger struct{}

func (c customLogger) LogRequest(method string, dump []byte) {
	var re = regexp.MustCompile(`(<vatNumber>)[\s\S]*?(<\/vatNumber>)`)
	maskedResponse := re.ReplaceAllString(string(dump), `${1}XXX${2}`)

	log.Printf("%s request: %s", method, maskedResponse)
}

func (c customLogger) LogResponse(method string, dump []byte) {
	if method == "checkVat" {
		return
	}

	log.Printf("Response: %s", dump)
}

func TestClient_Call_WithCustomLogger(t *testing.T) {
	soap, err := SoapClientWithConfig("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl",
		nil,
		&Config{Dump: true, Logger: &customLogger{}},
	)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	var res *Response

	res, err = soap.CallByStruct(CheckVatRequest{
		CountryCode: "IE",
		VatNumber:   "6388047V",
	})
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rv)
	if rv.CountryCode != "IE" {
		t.Errorf("error: %+v", rv)
	}

	_, err = soap.CallByStruct(nil)
	if err == nil {
		t.Error("err can't be nil")
	}
}

func TestClient_CallByStruct(t *testing.T) {
	soap, err := SoapClient("http://ec.europa.eu/taxation_customs/vies/checkVatService.wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	var res *Response

	res, err = soap.CallByStruct(CheckVatRequest{
		CountryCode: "IE",
		VatNumber:   "6388047V",
	})
	if err != nil {
		t.Errorf("error in soap call: %s", err)
	}

	res.Unmarshal(&rv)
	if rv.CountryCode != "IE" {
		t.Errorf("error: %+v", rv)
	}

	_, err = soap.CallByStruct(nil)
	if err == nil {
		t.Error("err can't be nil")
	}
}

func TestClient_Call_NonUtf8(t *testing.T) {
	soap, err := SoapClient("https://demo.ilias.de/webservice/soap/server.php?wsdl", nil)
	if err != nil {
		t.Errorf("error not expected: %s", err)
	}

	_, err = soap.Call("login", Params{"client": "demo", "username": "robert", "password": "iliasdemo"})
	if err == nil {
		t.Errorf("err can't be nil")
	}
}

func TestProcess_doRequest(t *testing.T) {
	c := &process{
		Client: &Client{
			HTTPClient: &http.Client{},
		},
	}

	_, err := c.doRequest(context.Background(), "")
	if err == nil {
		t.Errorf("body is empty")
	}

	_, err = c.doRequest(context.Background(), "://teste.")
	if err == nil {
		t.Errorf("invalid WSDL")
	}

	_, err = c.doRequest(context.Background(), "https://google.com/non-existent-url")
	if err == nil {
		t.Errorf("err can't be nil")
	}

	doneC := make(chan struct{})
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-doneC
	}))
	defer ts.Close()

	ctx, cancelF := context.WithTimeout(context.Background(), time.Second)
	defer cancelF()
	_, err = c.doRequest(ctx, ts.URL)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("request didn't timeout")
	}
	close(doneC)
}
