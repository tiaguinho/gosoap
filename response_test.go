package gosoap

import (
	"encoding/xml"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUnmarshal(t *testing.T) {
	var testCases = []struct {
		description  string
		response     *Response
		decodeStruct interface{}
		isFaultError bool
	}{
		{
			description: "case: fault error",
			response: &Response{
				Body: []byte(`
				<soap:Fault>
					<faultcode>soap:Server</faultcode>
					<faultstring>Qube.Mama.SoapException: The remote server returned an error: (550) File unavailable (e.g., file not found, no access).
				The remote server returned an error: (550) File unavailable (e.g., file not found, no access).
					</faultstring>
					<detail>
					</detail>
			</soap:Fault>	
				`),
			},
			decodeStruct: &struct{}{},
			isFaultError: true,
		},
		{
			description: "case: unmarshal error",
			response: &Response{
				Body: []byte(`
					<GetJobsByIdsResponse
						xmlns="http://webservices.qubecinema.com/XP/Usher/2009-09-29/">
						<GetJobsByIdsResult>
							<JobInfo>
								<ID>9e7d58d9-6f62-43e3-b189-5b1b58eea629</ID>
								<Status>Completed</Status>
								<Progress>0</Progress>
								<VerificationProgress>0</VerificationProgress>
								<EstimatedCompletionTime>0</EstimatedCompletionTime>
							</JobInfo>
						</GetJobsByIdsResult>
					</GetJobsByIdsResponse>
			
				`),
			},
			decodeStruct: &struct {
				XMLName            xml.Name `xml:"GetJobsByIsResponse"`
				GetJobsByIDsResult string
			}{},
			isFaultError: false,
		},
		{
			description: "case: nil error",
			response: &Response{
				Body: []byte(`
					<GetJobsByIdsResponse
						xmlns="http://webservices.qubecinema.com/XP/Usher/2009-09-29/">
						<GetJobsByIdsResult>
							<JobInfo>
								<ID>9e7d58d9-6f62-43e3-b189-5b1b58eea629</ID>
								<Status>Completed</Status>
								<Progress>0</Progress>
								<VerificationProgress>0</VerificationProgress>
								<EstimatedCompletionTime>0</EstimatedCompletionTime>
							</JobInfo>
						</GetJobsByIdsResult>
					</GetJobsByIdsResponse>
			
				`),
			},
			decodeStruct: &struct {
				XMLName            xml.Name `xml:"GetJobsByIdsResponse"`
				GetJobsByIDsResult string
			}{},
			isFaultError: false,
		},
		{
			description: "case: body is empty",
			response: &Response{
				Body: []byte(``),
			},
			decodeStruct: &struct {
				XMLName            xml.Name `xml:"GetJobsByIdsResponse"`
				GetJobsByIDsResult string
			}{},
			isFaultError: false,
		},
	}

	for _, testCase := range testCases {
		t.Logf("running %v test case", testCase.description)

		err := testCase.response.Unmarshal(testCase.decodeStruct)
		assert.Equal(t, testCase.isFaultError, IsFault(err))
	}
}

func TestIsFault(t *testing.T) {
	var testCases = []struct {
		description          string
		err                  error
		expectedIsFaultError bool
	}{
		{
			description: "case: fault error",
			err: FaultError{
				fault: &Fault{
					Code: "SOAP-ENV:Client",
				},
			},
			expectedIsFaultError: true,
		},
		{
			description:          "case: unmarshal error",
			err:                  fmt.Errorf("unmarshall err: .."),
			expectedIsFaultError: false,
		},
		{
			description:          "case: nil error",
			err:                  nil,
			expectedIsFaultError: false,
		},
	}

	for _, testCase := range testCases {
		t.Logf("running %v test case", testCase.description)

		isFaultErr := IsFault(testCase.err)
		assert.Equal(t, testCase.expectedIsFaultError, isFaultErr)
	}
}
