package gosoap

import (
	"fmt"
)

// Request Soap Request
type Request struct {
	Method string
	Params Params
}

// NewRequest creates a new soap request
func NewRequest(m string, p Params) *Request {
	return &Request{
		Method: m,
		Params: p,
	}
}

// RequestStruct soap request interface
type RequestStruct interface {
	SoapBuildRequest() *Request
}

// NewRequestByStruct create a new request using builder
func NewRequestByStruct(s RequestStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequest(), nil
}
