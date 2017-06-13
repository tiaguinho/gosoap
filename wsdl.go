package gosoap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type WsdlDefinitions struct {
	Name            string           `xml:"name,attr"`
	TargetNamespace string           `xml:"targetNamespace,attr"`
	Imports         []*WsdlImport    `xml:"http://schemas.xmlsoap.org/wsdl/ import"`
	Types           []*WsdlTypes     `xml:"http://schemas.xmlsoap.org/wsdl/ types"`
	Messages        []*WsdlMessage   `xml:"http://schemas.xmlsoap.org/wsdl/ message"`
	PortTypes       []*WsdlPortTypes `xml:"http://schemas.xmlsoap.org/wsdl/ portType"`
	Services        []*WsdlService   `xml:"http://schemas.xmlsoap.org/wsdl/ service"`
	Bindings        []*WsdlBinding   `xml:"http://schemas.xmlsoap.org/wsdl/ binding"`
}

type WsdlBinding struct {
	Name         string           `xml:"name,attr"`
	Type         string           `xml:"type,attr"`
	Operations   []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
	SoapBindings []*SoapBinding   `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
}

type SoapBinding struct {
	Transport string `xml:"transport,attr"`
}

type WsdlTypes struct {
	XsdSchema []*XsdSchema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type WsdlImport struct {
	Namespace string `xml:"namespace,attr"`
	Location  string `xml:"location,attr"`
}

type WsdlMessage struct {
	Name  string             `xml:"name,attr"`
	Parts []*WsdlMessagePart `xml:"http://schemas.xmlsoap.org/wsdl/ part"`
}

type WsdlMessagePart struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type WsdlPortTypes struct {
	Name       string           `xml:"name,attr"`
	Operations []*WsdlOperation `xml:"http://schemas.xmlsoap.org/wsdl/ operation"`
}

type WsdlOperation struct {
	Name           string                 `xml:"name,attr"`
	Inputs         []*WsdlOperationInput  `xml:"http://schemas.xmlsoap.org/wsdl/ input"`
	Outputs        []*WsdlOperationOutput `xml:"http://schemas.xmlsoap.org/wsdl/ output"`
	Faults         []*WsdlOperationFault  `xml:"http://schemas.xmlsoap.org/wsdl/ fault"`
	SoapOperations []*SoapOperation       `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
}

type WsdlOperationInput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationOutput struct {
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlOperationFault struct {
	Name       string `xml:"name,attr"`
	Message    string `xml:"message,attr"`
	WsawAction string `xml:"http://www.w3.org/2006/05/addressing/wsdl Action,attr"`
}

type WsdlService struct {
	Name  string      `xml:"name,attr"`
	Ports []*WsdlPort `xml:"http://schemas.xmlsoap.org/wsdl/ port"`
}

type WsdlPort struct {
	Name          string         `xml:"name,attr"`
	Binding       string         `xml:"binding,attr"`
	SoapAddresses []*SoapAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type SoapAddress struct {
	Location string `xml:"location,attr"`
}

type SoapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
	Style      string `xml:"style,attr"`
}

type XsdSchema struct {
	TargetNamespace    string            `xml:"targetNamespace,attr"`
	ElementFormDefault string            `xml:"elementFormDefault,attr"`
	Imports            []*XsdImport      `xml:"http://www.w3.org/2001/XMLSchema import"`
	Elements           []*XsdElement     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []*XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type XsdImport struct {
	SchemaLocation string `xml:"schemaLocation,attr"`
	Namespace      string `xml:"namespace,attr"`
}

type XsdElement struct {
	Name        string          `xml:"name,attr"`
	Nillable    bool            `xml:"nillable,attr"`
	Type        string          `xml:"type,attr"`
	MinOccurs   string          `xml:"minOccurs,attr"`
	MaxOccurs   string          `xml:"maxOccurs,attr"`
	ComplexType *XsdComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleType  *XsdSimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type XsdComplexType struct {
	Name     string       `xml:"name,attr"`
	Sequence *XsdSequence `xml:"http://www.w3.org/2001/XMLSchema sequence"`
}

type XsdSimpleType struct {
	Name     string          `xml:"name,attr"`
	Sequence *XsdRestriction `xml:"http://www.w3.org/2001/XMLSchema restriction"`
}

type XsdSequence struct {
	Elements []*XsdElement `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type XsdRestriction struct {
	Base         string           `xml:"base,attr"`
	Pattern      *XsdPattern      `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	MinInclusive *XsdMinInclusive `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
	MaxInclusive *XsdMaxInclusive `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type XsdPattern struct {
	Value string `xml:"value,attr"`
}

type XsdMinInclusive struct {
	Value string `xml:"value,attr"`
}

type XsdMaxInclusive struct {
	Value string `xml:"value,attr"`
}

// getWsdlDefinitions sent request to the wsdl url and set definitions on struct
func getWsdlDefinitions(u string) (wsdl *WsdlDefinitions, err error) {
	r, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	err = xml.Unmarshal(b, &wsdl)

	return wsdl, err
}
