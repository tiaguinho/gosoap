package gosoap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type wsdlDefinitions struct {
	Name            string           `xml:"name,attr"`
	TargetNamespace string           `xml:"targetNamespace,attr"`
	Imports         []*wsdlImport    `xml:"http://schemas.xmlsoap.org/wsdl/ import"`
	Types           []*wsdlTypes     `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []*wsdlMessage   `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []*wsdlPortTypes `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Services        []*wsdlService   `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	Bindings        []*wsdlBinding   `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
}

type wsdlBinding struct {
	Name         string           `xml:"name,attr"`
	Type         string           `xml:"type,attr"`
	Operations   []*wsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
	SoapBindings []*soapBinding   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
}

type soapBinding struct {
	Transport string `xml:"transport,attr"`
}

type wsdlTypes struct {
	XsdSchema []*xsdSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type wsdlImport struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

type wsdlMessage struct {
	Name  string             `xml:"name,attr"`
	Parts []*wsdlMessagePart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type wsdlMessagePart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type wsdlPortTypes struct {
	Name       string           `xml:"name,attr"`
	Operations []*wsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type wsdlOperation struct {
	Name           string                 `xml:"name,attr"`
	Inputs         []*wsdlOperationInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Outputs        []*wsdlOperationOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	Faults         []*wsdlOperationFault  `xml:"http://schemas.xmlsoap.org/wsdl/ fault"`
	SoapOperations []*soapOperation       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
}

type wsdlOperationInput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type wsdlOperationOutput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type wsdlOperationFault struct {
	Name       string `xml:"name,attr"`
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type wsdlService struct {
	Name  string      `xml:"name,attr"`
	Ports []*wsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type wsdlPort struct {
	Name          string         `xml:"name,attr"`
	Binding       string         `xml:"binding,attr"`
	SoapAddresses []*soapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type soapAddress struct {
	Location string `xml:"location,attr"`
}

type soapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

type xsdSchema struct {
	TargetNamespace    string            `xml:"targetNamespace,attr"`
	ElementFormDefault string            `xml:"elementFormDefault,attr"`
	Imports            []*xsdImport      `xml:"http://www.w3.org/2001/XMLSchema import"`
	Elements           []*xsdElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []*xsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type xsdImport struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
	Namespace      string `xml:"namespace,attr"`
}

type xsdElement struct {
	Name        string          `xml:"name,attr"`
	Nillable    bool            `xml:"nillable,attr"`
	Type        string          `xml:"type,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	ComplexType *xsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleType  *xsdSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type xsdComplexType struct {
	Name     string       `xml:"name,attr"`
	Sequence *xsdSequence `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type xsdSimpleType struct {
	Name     string          `xml:"name,attr"`
	Sequence *xsdRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type xsdSequence struct {
	Elements []*xsdElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type xsdRestriction struct {
	Base         string           `xml:"base,attr"`
	Pattern      *xsdPattern      `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	MinInclusive *xsdMinInclusive `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
	MaxInclusive *xsdMaxInclusive `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type xsdPattern struct {
	Value string `xml:"value,attr"`
}

type xsdMinInclusive struct {
	Value string `xml:"value,attr"`
}

type xsdMaxInclusive struct {
	Value string `xml:"value,attr"`
}

// getwsdlDefinitions sent request to the wsdl url and set definitions on struct
func getWsdlDefinitions(u string) (wsdl *wsdlDefinitions, err error) {
	r, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	err = xml.Unmarshal(b, &wsdl)

	return wsdl, err
}

// Fault response
type Fault struct {
	Code        string `xml:"faultcode"`
	Description string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}
