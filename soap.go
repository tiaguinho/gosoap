package gosoap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"moul.io/http2curl"
)

// HeaderParams holds params specific to the header
type HeaderParams map[string]string

// Params type is used to set the params in soap request
type Params map[string]interface{}

// SoapClient return new *Client to handle the requests with the WSDL
func SoapClient(wsdl string) (*Client, error) {
	_, err := url.Parse(wsdl)
	if err != nil {
		return nil, err
	}

	d, err := getWsdlDefinitions(wsdl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		WSDL:        wsdl,
		URL:         strings.TrimSuffix(d.TargetNamespace, "/"),
		Definitions: d,
	}

	return c, nil
}

// Client struct hold all the informations about WSDL,
// request and response of the server
type Client struct {
	HttpClient   *http.Client
	WSDL         string
	URL          string
	Method       string
	SoapAction   string
	Params       Params
	HeaderName   string
	HeaderParams HeaderParams
	Definitions  *wsdlDefinitions
	Body         []byte
	Header       []byte
	Username     string
	Password     string

	payload []byte
}

// GetLastRequest returns the last request
func (c *Client) GetLastRequest() []byte {
	return c.payload
}

// Call call's the method m with Params p
func (c *Client) Call(m string, p Params) (err error) {
	c.Method = m
	c.Params = p
	c.SoapAction = c.Definitions.GetSoapActionFromWsdlOperation(c.Method)
	if c.SoapAction == "" {
		c.SoapAction = fmt.Sprintf("%s/%s", c.URL, c.Method)
	}

	c.payload, err = xml.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	b, err := c.doRequest(c.Definitions.Services[0].Ports[0].SoapAddresses[0].Location)
	if err != nil {
		return err
	}

	var soap SoapEnvelope
	err = xml.Unmarshal(b, &soap)

	c.Body = soap.Body.Contents
	c.Header = soap.Header.Contents

	return err
}

// Unmarshal get the body and unmarshal into the interface
func (c *Client) Unmarshal(v interface{}) error {
	if len(c.Body) == 0 {
		return fmt.Errorf("Body is empty")
	}

	var f Fault
	xml.Unmarshal(c.Body, &f)
	if f.Code != "" {
		return fmt.Errorf("[%s]: %s", f.Code, f.Description)
	}

	return xml.Unmarshal(c.Body, v)
}

// doRequest makes new request to the server using the c.Method, c.URL and the body.
// body is enveloped in Call method
func (c *Client) doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	if c.Username != "" && c.Password != "" {
		req.SetBasicAuth(c.Username, c.Password)
	}

	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}
	}

	req.ContentLength = int64(len(c.payload))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("SOAPAction", c.SoapAction)

	command, _ := http2curl.GetCurlCommand(req)
	logrus.Debugf("\nHTTP REQUEST AS cURL:\n%+v\n\n", command)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// SoapEnvelope struct
type SoapEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Header  SoapHeader
	Body    SoapBody
}

// SoapHeader struct
type SoapHeader struct {
	XMLName  struct{} `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

// SoapBody struct
type SoapBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}
