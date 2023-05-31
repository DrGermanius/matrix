package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func init() {
	mux := NewHandler()
	go http.ListenAndServe(":8081", mux)
}

const (
	unexpectedBody = "handler returned unexpected body: got %v want %v"
	unexpectedCode = "handler returned unexpected code: got %v want %v"

	validPath          = "testData/matrix.csv"
	wrongExtensionPath = "testData/matrix.txt"
	invalidCSVPath     = "testData/invalidCSV.csv"
	emptyPath          = "testData/emptyFIle.csv"
	notSquarePath      = "testData/notsquare.csv"

	defaultURL = "http://localhost:8081"
)

func TestHandler_Echo(t *testing.T) {
	url := fmt.Sprintf("%s%s", defaultURL, "/echo")
	okReq, writer := SetupRequest(validPath, url, t)
	okReq.Header.Set("Content-Type", writer.FormDataContentType())

	wrongFormatReq, writer := SetupRequest(wrongExtensionPath, url, t)
	wrongFormatReq.Header.Set("Content-Type", writer.FormDataContentType())

	invalidCsvReq, writer := SetupRequest(invalidCSVPath, url, t)
	invalidCsvReq.Header.Set("Content-Type", writer.FormDataContentType())

	emptyFileReq, writer := SetupRequest(emptyPath, url, t)
	emptyFileReq.Header.Set("Content-Type", writer.FormDataContentType())

	notSquareReq, writer := SetupRequest(notSquarePath, url, t)
	notSquareReq.Header.Set("Content-Type", writer.FormDataContentType())

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "echo endpoint happy path",
			when: "everything is OK",
			then: "same matrix should be returned",

			args:     args{req: okReq},
			wantBody: "1,2,3\n4,5,6\n7,8,9\n",
			wantCode: http.StatusOK,
		},
		{
			name: "echo endpoint unhappy path with wrong format",
			when: "wrong file format sent",
			then: "error should be returned",

			args:     args{req: wrongFormatReq},
			wantBody: "invalid file format, only CSV allowed\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "echo endpoint unhappy path with invalid csv",
			when: "invalid csv file sent",
			then: "error should be returned",

			args:     args{req: invalidCsvReq},
			wantBody: "record on line 2: wrong number of fields\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "echo endpoint unhappy path with empty file",
			when: "request sent without file",
			then: "error should be returned",

			args:     args{req: emptyFileReq},
			wantBody: "matrix shouldn't be empty\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "echo endpoint unhappy path with not square matrix",
			when: "the sent matrix is not square",
			then: "error should be returned",

			args:     args{req: notSquareReq},
			wantBody: "matrix should be square\n",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(testCase.args.req)
			if err != nil {
				t.Error(err)
			}
			if status := resp.StatusCode; status != testCase.wantCode {
				t.Errorf(unexpectedCode, status, testCase.wantCode)
			}

			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if string(resBody) != testCase.wantBody {
				t.Errorf(unexpectedBody, string(resBody), testCase.wantBody)
			}
		})
	}
}

func TestHandler_Invert(t *testing.T) {
	url := fmt.Sprintf("%s%s", defaultURL, "/invert")
	okReq, writer := SetupRequest(validPath, url, t)
	okReq.Header.Set("Content-Type", writer.FormDataContentType())

	wrongFormatReq, writer := SetupRequest(wrongExtensionPath, url, t)
	wrongFormatReq.Header.Set("Content-Type", writer.FormDataContentType())

	invalidCsvReq, writer := SetupRequest(invalidCSVPath, url, t)
	invalidCsvReq.Header.Set("Content-Type", writer.FormDataContentType())

	emptyFileReq, writer := SetupRequest(emptyPath, url, t)
	emptyFileReq.Header.Set("Content-Type", writer.FormDataContentType())

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "invert endpoint happy path",
			when: "everything is OK",
			then: "inverted matrix should be returned",

			args:     args{req: okReq},
			wantBody: "1,4,7\n2,5,8\n3,6,9\n",
			wantCode: http.StatusOK,
		},
		{
			name: "invert endpoint unhappy path with wrong format",
			when: "wrong file format sent",
			then: "error should be returned",

			args:     args{req: wrongFormatReq},
			wantBody: "invalid file format, only CSV allowed\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "invert endpoint unhappy path with invalid csv",
			when: "invalid csv file sent",
			then: "error should be returned",

			args:     args{req: invalidCsvReq},
			wantBody: "record on line 2: wrong number of fields\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "invert endpoint unhappy path with empty file",
			when: "request sent without file",
			then: "error should be returned",

			args:     args{req: emptyFileReq},
			wantBody: "matrix shouldn't be empty\n",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(testCase.args.req)
			if err != nil {
				t.Error(err)
			}
			if status := resp.StatusCode; status != testCase.wantCode {
				t.Errorf(unexpectedCode, status, testCase.wantCode)
			}

			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if string(resBody) != testCase.wantBody {
				t.Errorf(unexpectedBody, string(resBody), testCase.wantBody)
			}
		})
	}
}

func TestHandler_Flatten(t *testing.T) {
	url := fmt.Sprintf("%s%s", defaultURL, "/flatten")
	okReq, writer := SetupRequest(validPath, url, t)
	okReq.Header.Set("Content-Type", writer.FormDataContentType())

	wrongFormatReq, writer := SetupRequest(wrongExtensionPath, url, t)
	wrongFormatReq.Header.Set("Content-Type", writer.FormDataContentType())

	invalidCsvReq, writer := SetupRequest(invalidCSVPath, url, t)
	invalidCsvReq.Header.Set("Content-Type", writer.FormDataContentType())

	emptyFileReq, writer := SetupRequest(emptyPath, url, t)
	emptyFileReq.Header.Set("Content-Type", writer.FormDataContentType())

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "flatten endpoint happy path",
			when: "everything is OK",
			then: "flatten matrix should be returned",

			args:     args{req: okReq},
			wantBody: "1,2,3,4,5,6,7,8,9",
			wantCode: http.StatusOK,
		},
		{
			name: "flatten endpoint unhappy path with wrong format",
			when: "wrong file format sent",
			then: "error should be returned",

			args:     args{req: wrongFormatReq},
			wantBody: "invalid file format, only CSV allowed\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flatten endpoint unhappy path with invalid csv",
			when: "invalid csv file sent",
			then: "error should be returned",

			args:     args{req: invalidCsvReq},
			wantBody: "record on line 2: wrong number of fields\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "flatten endpoint unhappy path with empty file",
			when: "request sent without file",
			then: "error should be returned",

			args:     args{req: emptyFileReq},
			wantBody: "matrix shouldn't be empty\n",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(testCase.args.req)
			if err != nil {
				t.Error(err)
			}
			if status := resp.StatusCode; status != testCase.wantCode {
				t.Errorf(unexpectedCode, status, testCase.wantCode)
			}

			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if string(resBody) != testCase.wantBody {
				t.Errorf(unexpectedBody, string(resBody), testCase.wantBody)
			}
		})
	}
}

func TestHandler_Sum(t *testing.T) {
	url := fmt.Sprintf("%s%s", defaultURL, "/sum")
	okReq, writer := SetupRequest(validPath, url, t)
	okReq.Header.Set("Content-Type", writer.FormDataContentType())

	wrongFormatReq, writer := SetupRequest(wrongExtensionPath, url, t)
	wrongFormatReq.Header.Set("Content-Type", writer.FormDataContentType())

	invalidCsvReq, writer := SetupRequest(invalidCSVPath, url, t)
	invalidCsvReq.Header.Set("Content-Type", writer.FormDataContentType())

	emptyFileReq, writer := SetupRequest(emptyPath, url, t)
	emptyFileReq.Header.Set("Content-Type", writer.FormDataContentType())

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "sum endpoint happy path",
			when: "everything is OK",
			then: "sum of all matrix element should be returned",

			args:     args{req: okReq},
			wantBody: "45",
			wantCode: http.StatusOK,
		},
		{
			name: "sum endpoint unhappy path with wrong format",
			when: "wrong file format sent",
			then: "error should be returned",

			args:     args{req: wrongFormatReq},
			wantBody: "invalid file format, only CSV allowed\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "sum endpoint unhappy path with invalid csv",
			when: "invalid csv file sent",
			then: "error should be returned",

			args:     args{req: invalidCsvReq},
			wantBody: "record on line 2: wrong number of fields\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "sum endpoint unhappy path with empty file",
			when: "request sent without file",
			then: "error should be returned",

			args:     args{req: emptyFileReq},
			wantBody: "matrix shouldn't be empty\n",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(testCase.args.req)
			if err != nil {
				t.Error(err)
			}
			if status := resp.StatusCode; status != testCase.wantCode {
				t.Errorf(unexpectedCode, status, testCase.wantCode)
			}

			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if string(resBody) != testCase.wantBody {
				t.Errorf(unexpectedBody, string(resBody), testCase.wantBody)
			}
		})
	}
}

func TestHandler_Multiply(t *testing.T) {
	url := fmt.Sprintf("%s%s", defaultURL, "/multiply")
	okReq, writer := SetupRequest(validPath, url, t)
	okReq.Header.Set("Content-Type", writer.FormDataContentType())

	wrongFormatReq, writer := SetupRequest(wrongExtensionPath, url, t)
	wrongFormatReq.Header.Set("Content-Type", writer.FormDataContentType())

	invalidCsvReq, writer := SetupRequest(invalidCSVPath, url, t)
	invalidCsvReq.Header.Set("Content-Type", writer.FormDataContentType())

	emptyFileReq, writer := SetupRequest(emptyPath, url, t)
	emptyFileReq.Header.Set("Content-Type", writer.FormDataContentType())

	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args     args
		wantBody string
		wantCode int
	}{
		{
			name: "multiply endpoint happy path",
			when: "everything is OK",
			then: "multiplication of all matrix element should be returned",

			args:     args{req: okReq},
			wantBody: "362880",
			wantCode: http.StatusOK,
		},
		{
			name: "multiply endpoint unhappy path with wrong format",
			when: "wrong file format sent",
			then: "error should be returned",

			args:     args{req: wrongFormatReq},
			wantBody: "invalid file format, only CSV allowed\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "multiply endpoint unhappy path with invalid csv",
			when: "invalid csv file sent",
			then: "error should be returned",

			args:     args{req: invalidCsvReq},
			wantBody: "record on line 2: wrong number of fields\n",
			wantCode: http.StatusBadRequest,
		},
		{
			name: "multiply endpoint unhappy path with empty file",
			when: "request sent without file",
			then: "error should be returned",

			args:     args{req: emptyFileReq},
			wantBody: "matrix shouldn't be empty\n",
			wantCode: http.StatusBadRequest,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(testCase.args.req)
			if err != nil {
				t.Error(err)
			}
			if status := resp.StatusCode; status != testCase.wantCode {
				t.Errorf(unexpectedCode, status, testCase.wantCode)
			}

			resBody, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}
			if string(resBody) != testCase.wantBody {
				t.Errorf(unexpectedBody, string(resBody), testCase.wantBody)
			}
		})
	}
}

func SetupRequest(filePath string, url string, t *testing.T) (*http.Request, *multipart.Writer) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	formFile, err := writer.CreateFormFile(multipartFileKey, filePath)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	// use io.Copy for copying data with headers
	_, err = io.Copy(formFile, file)
	if err != nil {
		t.Fatal(err)
	}

	// two lines below shouldn't be deferred by the reason of building multipart request
	file.Close()
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		t.Fatal(err)
	}

	return req, writer
}
