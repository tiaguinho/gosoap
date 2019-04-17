# Go Soap [![Build Status](https://travis-ci.org/tiaguinho/gosoap.svg?branch=master)](https://travis-ci.org/tiaguinho/gosoap) [![GoDoc](https://godoc.org/github.com/tiaguinho/gosoap?status.png)](https://godoc.org/github.com/tiaguinho/gosoap) [![Go Report Card](https://goreportcard.com/badge/github.com/tiaguinho/gosoap)](https://goreportcard.com/report/github.com/tiaguinho/gosoap) [![codecov](https://codecov.io/gh/tiaguinho/gosoap/branch/master/graph/badge.svg)](https://codecov.io/gh/tiaguinho/gosoap) [![patreon](https://img.shields.io/badge/patreon-donate-yellow.svg)](https://www.patreon.com/temporin)
package to help with SOAP integrations (client)

### Install

```bash
go get github.com/tiaguinho/gosoap
```

### Example

```go
package main

import (
	"github.com/tiaguinho/gosoap"
	"fmt"
)

type GetGeoIPResponse struct {
	GetGeoIPResult GetGeoIPResult
}

type GetGeoIPResult struct {
	ReturnCode        string
	IP                string
	ReturnCodeDetails string
	CountryName       string
	CountryCode       string
}

func main() {
	soap, err := gosoap.SoapClient("http://www.webservicex.net/geoipservice.asmx?WSDL")
	if err != nil {
		fmt.Errorf("error not expected: %s", err)
	}

	params := gosoap.Params{
		"IPAddress": "8.8.8.8",
	}

	res, err := soap.Call("GetGeoIP", params)
	if err != nil {
		fmt.Errorf("error in soap call: %s", err)
	}

	r := GetGeoIPResponse{}

	res.Unmarshal(&r)
	if r.GetGeoIPResult.CountryCode != "USA" {
		fmt.Errorf("error: %+v", r)
	}
}
```
