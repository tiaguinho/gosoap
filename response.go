package gosoap

import (
	"encoding/xml"
	"fmt"
)

type Response struct {
	Body    []byte
	Header  []byte
	Payload []byte
}

// Unmarshal get the body and unmarshal into the interface
func (r *Response) Unmarshal(v interface{}) error {
	if len(r.Body) == 0 {
		return fmt.Errorf("Body is empty")
	}

	var f Fault
	xml.Unmarshal(r.Body, &f)
	if f.Code != "" {
		return fmt.Errorf("[%s]: %s", f.Code, f.Description)
	}

	return xml.Unmarshal(r.Body, v)
}
