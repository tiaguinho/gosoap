package gosoap

import (
	"fmt"
)

type Request struct {
	Method string
	Params Params
}

func NewRequest(m string, p Params) *Request {
	return &Request{
		Method: m,
		Params: p,
	}
}

type RequestStruct interface {
	SoapBuildRequest() *Request
}

func NewRequestByStruct(s RequestStruct) (*Request, error) {
	if s == nil {
		return nil, fmt.Errorf("'s' cannot be 'nil'")
	}

	return s.SoapBuildRequest(), nil
}
