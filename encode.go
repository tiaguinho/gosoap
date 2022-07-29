package gosoap

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

var (
	soapPrefix                            = "soap"
	customEnvelopeAttrs map[string]string = nil
)

// SetCustomEnvelope define customizated envelope
func SetCustomEnvelope(prefix string, attrs map[string]string) {
	soapPrefix = prefix
	if attrs != nil {
		customEnvelopeAttrs = attrs
	}
}

// MarshalXML envelope the body and encode to xml
func (c process) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	tokens := &tokenData{}

	//start envelope
	if c.Client.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	namespace := ""
	if c.Client.Definitions.Types != nil {
		schema := c.Client.Definitions.Types[0].XsdSchema[0]
		namespace = schema.TargetNamespace
		if namespace == "" && len(schema.Imports) > 0 {
			namespace = schema.Imports[0].Namespace
		}
	}

	tokens.startEnvelope()
	if c.Client.HeaderParams != nil {
		tokens.startHeader(c.Client.HeaderName, namespace)
		if err := tokens.recursiveEncode(c.Client.HeaderParams); err != nil {
			return err
		}
		tokens.endHeader(c.Client.HeaderName)
	}

	err := tokens.startSoapBody(c.Request.Method, namespace)
	if err != nil {
		return err
	}

	err = tokens.bodyContents(c, namespace, e)
	if err != nil {
		return err
	}

	//end envelope
	tokens.endSoapBody()
	tokens.endEnvelope()

	if err := tokens.flush(e); err != nil {
		return err
	}

	return e.Flush()
}

func (tokens *tokenData) bodyContents(c process, namespace string, e *xml.Encoder) error {
	isStruct := false
	t := reflect.TypeOf(c.Request.Params)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() == reflect.Struct {
		isStruct = true
	}
	if isStruct {
		// Just use encoding/xml directly for structs, which allows for much more
		// sophisticated control over the XML structure. The top-level element
		// is intrinsically part of the encoding, so we don't need to do that
		// for ourselves

		// Flush any pending tokens before we send things directly to the encoder
		if err := tokens.flush(e); err != nil {
			return err
		}
		if err := e.Encode(c.Request.Params); err != nil {
			return err
		}
		// if err := tokens.flush(e); err != nil {
		// 	return err
		// }
	} else {
		// For non-structs, we have to explicitly wrap a top-level element around
		// the actual data
		tokens.startBodyContents(c.Request.Method, namespace)
		if err := tokens.recursiveEncode(c.Request.Params); err != nil {
			return err
		}
		tokens.endBodyContents(c.Request.Method)
	}
	return nil
}

func (tokens *tokenData) flush(e *xml.Encoder) error {
	for _, t := range tokens.data {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}
	tokens.data = []xml.Token{}
	return nil
}

type tokenData struct {
	data []xml.Token
}

func (tokens *tokenData) recursiveEncode(hm interface{}) error {
	v := reflect.ValueOf(hm)

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: key.String(),
				},
			}

			tokens.data = append(tokens.data, t)
			if err := tokens.recursiveEncode(v.MapIndex(key).Interface()); err != nil {
				return err
			}
			tokens.data = append(tokens.data, xml.EndElement{Name: t.Name})
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := tokens.recursiveEncode(v.Index(i).Interface()); err != nil {
				return err
			}
		}
	case reflect.Array:
		if v.Len() == 2 {
			label := v.Index(0).Interface()
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: label.(string),
				},
			}

			tokens.data = append(tokens.data, t)
			if err := tokens.recursiveEncode(v.Index(1).Interface()); err != nil {
				return err
			}
			tokens.data = append(tokens.data, xml.EndElement{Name: t.Name})
		}
	case reflect.String:
		content := xml.CharData(v.String())
		tokens.data = append(tokens.data, content)
	case reflect.Struct:
		tokens.data = append(tokens.data, v.Interface())
	}
	return nil
}

func (tokens *tokenData) startEnvelope() {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Envelope", soapPrefix),
		},
	}

	if customEnvelopeAttrs == nil {
		e.Attr = []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
			{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		}
	} else {
		e.Attr = make([]xml.Attr, 0)
		for local, value := range customEnvelopeAttrs {
			e.Attr = append(e.Attr, xml.Attr{
				Name:  xml.Name{Space: "", Local: local},
				Value: value,
			})
		}
	}

	tokens.data = append(tokens.data, e)
}

func (tokens *tokenData) endEnvelope() {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Envelope", soapPrefix),
		},
	}

	tokens.data = append(tokens.data, e)
}

func (tokens *tokenData) startHeader(m, n string) {
	h := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Header", soapPrefix),
		},
	}

	if m == "" || n == "" {
		tokens.data = append(tokens.data, h)
		return
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens.data = append(tokens.data, h, r)
}

func (tokens *tokenData) endHeader(m string) {
	h := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Header", soapPrefix),
		},
	}

	if m == "" {
		tokens.data = append(tokens.data, h)
		return
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens.data = append(tokens.data, r, h)
}

func (tokens *tokenData) startSoapBody(m, n string) error {
	if m == "" || n == "" {
		return fmt.Errorf("method or namespace is empty")
	}

	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Body", soapPrefix),
		},
	}
	tokens.data = append(tokens.data, b)
	return nil
}

func (tokens *tokenData) startBodyContents(m, n string) {
	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}
	tokens.data = append(tokens.data, r)
}

// endToken close body of the envelope
func (tokens *tokenData) endBodyContents(m string) {
	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}
	tokens.data = append(tokens.data, r)
}

func (tokens *tokenData) endSoapBody() {
	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: fmt.Sprintf("%s:Body", soapPrefix),
		},
	}
	tokens.data = append(tokens.data, b)
}
