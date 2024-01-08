package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	httpmock "github.com/miknonny/http-client/mock"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases for package 'example'")

	// httpmock.StartMockServer()

	os.Exit(m.Run())
}

// TODO. rewrite the test in this file.
// Look into the mock server and write a better mock using httpTest also mock delayed request.
func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {

		//Initialization:
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoint"),
		})

		// Execution:
		endpoints, err := GetEndPoints()

		// Assertion.
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoint" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {

		// Initialization:
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url":123}`,
		})

		// Execution:
		endpoints, err := GetEndPoints()

		// Assertion.

		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		//fmt.Println(err.Error())
		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {

		// Initialization:
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url":"https://api.github.com/user"}`,
		})

		// Execution:
		endpoints, err := GetEndPoints()

		// Assertion.
		// at this point we should not run any more test if the first test fails.
		if err != nil {
			t.Fatalf("no error was expected %s", err)
		}

		if endpoints == nil {
			t.Fatalf("endpoints expected")
		}

		if endpoints.CurrentUser != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
