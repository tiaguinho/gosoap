package gosoap

import (
	"encoding/xml"
)

var tokens []xml.Token

//
func (c Client) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	if len(c.Params) == 0 {
		return nil
	}

	tokens = []xml.Token{}

	//start envelope
	startToken(c.Method, c.Definitions.TargetNamespace)
	for k, v := range c.Params {
		t := xml.StartElement{
			Name: xml.Name{
				Space: "",
				Local: k,
			},
		}

		tokens = append(tokens, t, xml.CharData(v), xml.EndElement{t.Name})
	}
	//end envelope
	endToken(c.Method)

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	err := e.Flush()
	if err != nil {
		return err
	}

	return nil
}

//
func startToken(m, n string) {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
		Attr: []xml.Attr{
			{xml.Name{"", "xmlns:xsi"}, "http://www.w3.org/2001/XMLSchema-instance"},
			{xml.Name{"", "xmlns:xsd"}, "http://www.w3.org/2001/XMLSchema"},
			{xml.Name{"", "xmlns:soap"}, "http://schemas.xmlsoap.org/soap/envelope/"},
		},
	}

	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{xml.Name{"", "xmlns"}, n},
		},
	}

	tokens = append(tokens, e, b, r)
}

func endToken(m string) {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
	}

	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, b, e)
}
