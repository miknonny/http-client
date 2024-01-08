package examples

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	httpmock "github.com/miknonny/http-client/mock"
)

func TestCreateRepo(t *testing.T) {
	t.Run("timeoutFromGithub", func(t *testing.T) {
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			Error: errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(&repository)

		if repo != nil {
			t.Error("no repo expected when we get a timeout from github.")
		}

		if err == nil {
			t.Error("an error is expected whe we get an erro from github.")
		}
		fmt.Println(err.Error())
		if err.Error() != "timeout from github" {
			t.Error("invalid error message")
		}
	})

	t.Run("GithubPostError", func(t *testing.T) {
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusUnauthorized,
			ResponseBody:       `{"message": "Requires authentication","documentation_url": "https://docs.github.com/rest/repos/repos#create-a-repository-for-the-authenticated-user"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(&repository)

		if repo != nil {
			t.Fatal("was expecting repo to be nil")
		}

		if err == nil {
			t.Error("expected an error but did not get one.")
		}

		if !strings.Contains(err.Error(), "Requires authentication") {
			t.Error("invalid error message received")
		}

	})

	t.Run("UnmarshalGithubError", func(t *testing.T) {
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusUnauthorized,
			ResponseBody:       `{"message": "Requires authentication","documentation_url": "https://docs.github.com/rest/repos/repos#create-a-repository-for-the-authenticated-user}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(&repository)

		if repo != nil {
			t.Fatal("was expecting repo to be nil")
		}

		if err == nil {
			t.Error("expected an error but did not get one.")
		}

		if !strings.Contains(err.Error(), "error processing github error response when creating a new repo") {
			t.Error("invalid error message received")
		}

	})

	t.Run("UnmarshalRepositoryError", func(t *testing.T) {
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"name":"sweet}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(&repository)

		if repo != nil {
			t.Fatal("was expecting repo to be nil")
		}

		if err == nil {
			t.Error("expected an error but did not get one.")
		}

		if !strings.Contains(err.Error(), "unexpected end of JSON input") {
			t.Error("invalid error message received")
		}

	})

	t.Run("No Error", func(t *testing.T) {
		httpmock.DeleteAllMocks()
		httpmock.AddMock(&httpmock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","private":true}`,

			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":123,"name":"test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(&repository)

		if err != nil {
			t.Error("no error expected when we get a valid response")
		}

		if repo == nil {
			t.Fatal("a valid repo was expected")
		}

		if repo.Name != repository.Name {
			t.Error("invalid repository name")
		}
	})
}
