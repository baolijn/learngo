package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"fmt"
)

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}
func errorPanic(writer http.ResponseWriter, request *http.Request) error{
	panic(123)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error{
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error{
	return os.ErrNotExist
}

func errNoPermission(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter,
	_ *http.Request) error {
		fmt.Fprintln(writer, "no error")
	return nil
}

func TestErrWrapper(t *testing.T)  {
	tests := []struct{
		h appHandler
		code int
		message string
	}{
		{errorPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error"},
		{errNotFound, 404, "Not Found"},
		{errNoPermission, 403, "Forbidden"},
		{errUnknown, 500, "Internal Server Error"},
		{noError, 200, "no error"},
	}

	for _, tt := range tests{
		f := errorWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://baidu.com",
			nil,
			)
		f(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code || body != tt.message {
			t.Errorf("expect (%d, %s) " +
				"got (%d, %s)",
					tt.code, tt.message,
						response.Code, body)
		}
	}
}
