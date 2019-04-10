package gosoap

import (
	"fmt"
	"os"
	"testing"
)

func Test_getWsdlBody(t *testing.T) {
	type args struct {
		u string
	}
	dir, _ := os.Getwd()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				u: "http://[::1]:namedport",
			},
			wantErr: true,
		},
		{
			args: args{
				u: fmt.Sprintf("%s/%s", dir, "testdata/ipservice.wsdl"),
			},
			wantErr: true,
		},
		{
			args: args{
				u: fmt.Sprintf("file://%s/%s", dir, "testdata/ipservice.wsdl"),
			},
			wantErr: false,
		},
		{
			args: args{
				u: "file:",
			},
			wantErr: true,
		},
		{
			args: args{
				u: "https://www.google.com/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getWsdlBody(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("getwsdlBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
