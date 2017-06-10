package gosoap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Params map[string]string

//
func SoapClient(wsdl string) (*Client, error) {
	d, err := getWsdlDefinitions(wsdl)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(wsdl)
	if err != nil {
		return nil, err
	}

	c := &Client{
		WSDL:        wsdl,
		URL:         fmt.Sprintf("%s://%s", u.Scheme, u.Host),
		Definitions: d,
	}

	return c, nil
}

//
type Client struct {
	WSDL        string
	URL         string
	Definitions *WsdlDefinitions
	Method      string
	Params      Params
}

//
func (c *Client) Call(m string, p Params) ([]byte, error) {
	c.Method = m
	c.Params = p

	b, err := xml.MarshalIndent(c, "", "")
	if err != nil {
		panic(err)
	}

	b, err = c.doRequest(b)
	if err != nil {

	}

	var soap SoapEnvelope
	err = xml.Unmarshal(b, &soap)

	return soap.Body.Contents, err
}

//
func (c *Client) doRequest(body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.WSDL, bytes.NewBuffer(body))
	if err != nil {

	}

	req.ContentLength = int64(len(body))

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("Accept", "text/xml")
	req.Header.Add("SOAPAction", fmt.Sprintf("%s/%s", c.URL, c.Method))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	return b, err
}

//
type SoapEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}
